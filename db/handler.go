package db

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/constant/errors"
	"github.com/WayneShenHH/servermodule/logger"
)

type arangoHdr struct {
	db driver.Database
}

const (
	template = `
	FOR <doc> IN @@Collection
	FILTER IS_NULL(<doc>.migrate) || <doc>.migrate < @version
	UPDATE <doc> WITH {
	migrate: @version,
	<expr>
	} IN @@Collection
		`
	arrayTemplete = `
	<aname>:(
		FOR <item> IN <parrent>.<aname>
		<filter>
		RETURN IS_NULL(<item>.<next>) ? <item> : 
		MERGE(<item>, {<merge>}))`

	columnTemplate = `
	<cname>:(
		IS_NULL(<parrent>.<next>) ? <parrent> : 
		MERGE(<parrent>, {<merge>}))`

	arrayFilterTemplate = `FILTER <expr>`

	leafTemplate = `
	<cname>: <value>`

	leafArrayTemplate = `
	<aname>:(
		FOR <leaf> IN <parrent>
		RETURN <value>)`

	docItem            = "document"
	rootLevel          = "$root"
	arrayItemPrefix    = "a"
	arrayItemPostfix   = "_item"
	leafArrayItem      = "leaf"
	leftSquareBracket  = '['
	rightSquareBracket = ']'
	arraySelfItem      = "@"
	dot                = '.'
	questionMark       = '?'
	comma              = ","
)

// NewArango 建立 Arango 連線
func NewArango(cfg *config.DatabaseConfig) NoSQL {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{cfg.Host},
		ConnLimit: cfg.MaxConns,
	})
	if err != nil {
		logger.Fatalf(`connect to arangodb failed, endpoint: %v`, cfg.Host)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(cfg.Username, cfg.Password),
	})
	if err != nil {
		logger.Fatalf(`arangodb newclient failed, endpoint: %v`, cfg.Host)
	}

	ctx := context.Background()
	db, err := client.Database(ctx, cfg.Name)
	if err != nil {
		// handle error
		logger.Fatalf(`arangodb get Database failed, endpoint: %v, dbname: %v`, cfg.Host, cfg.Name)
	}

	return &arangoHdr{
		db: db,
	}
}

/** 更新特定 collection 的 document
 * 針對 Collection 操作
 *
 * @param collection 集合名稱
 * @param key document key
 * @param data document dao
 * @return int64 狀態碼 */
func (hdr *arangoHdr) UpdateDocument(collection, key string, data interface{}) int64 {
	// col, err := hdr.db.Collection(context.Background(), collection)
	// if err != nil {
	// 	logger.Debugf("db/arangoHdr/UpdateDocument: %v", err)
	// 	return errors.DatabaseCollection
	// }

	// meta, err := col.UpdateDocument(context.Background(), key, data)
	// if meta.ID.IsEmpty() {
	// 	logger.Debugf("db/arangoHdr/UpdateDocument: meta.ID is empty")
	// 	return errors.DatabaseMetaEmptyID
	// }
	// if err != nil {
	// 	logger.Debugf("db/arangoHdr/UpdateDocument: %v", err)
	// 	return errors.DatabaseUpdate
	// }

	jsonData, err := json.Marshal(data)
	if err != nil {
		logger.Debugf("db/arangoHdr/json.Marshal: %v", err)
		return errors.InputValueInvalid
	}

	action := `function (Params) {
		const db = require('@arangodb').db;
		const records = db._collection(Params[0]);
		records.update({_key:""+Params[1]}, JSON.parse(Params[2]));
		return 1}`

	options := &driver.TransactionOptions{
		MaxTransactionSize: 1000,
		WriteCollections:   []string{collection},
		ReadCollections:    []string{collection},
		Params:             []interface{}{collection, key, string(jsonData)},
		WaitForSync:        false,
	}

	_, err = hdr.db.Transaction(context.Background(), action, options)
	if err != nil {
		logger.Debugf("db/arangoHdr/UpdateDocument/Transaction: %v", err)
		return errors.DatabaseUpdate
	}

	return 0
}

/** NestedUpdate 更新特定 collection 的 document
 * 依照傳入的欄位定義(json path)，更新數值，如果有多個欄位位於同一層級，將合併成一個aql送出
 * 陣列篩選定義 :
 *  '?(expr)' : 只更新 expr 過濾出來的元素，若 expr 內含有欄位名稱定義，則自動替換成該欄位實際宣告的別名
 *        '*' : 全部更新
 *        '@' : 當前元素 例 processData[?(@.eventCode=1)]
 *    數值定義 : params 內每個欄位定義都有對應的數值定義，若含有跟欄位名稱相同的定義，則自動替換成該欄位實際宣告的別名
 *         例 : "history.paradise.MemberIncome": "MemberIncome * 100", 即該欄位更新成原本的一百倍
 * @param collection 集合名稱
 * @param version 設定 migrate 編號，更新成功時寫入 document.migrate
 * @param params 欄位定義
 * @return int64 狀態碼 */
func (hdr *arangoHdr) NestedUpdate(collection, version string, params map[string]string) int64 {
	bindVars := map[string]interface{}{
		"@Collection": collection,
		"version":     version,
	}

	for key, val := range groupValues(params) {
		levelofcolumns := getLevels(key)
		scripts := "<merge>" //一個欄位的所有層級更新語法
		ctrl := &parsectrl{
			parrent: docItem,
			value:   val,
			alias:   make(map[string]string),
		}
		for idx := range levelofcolumns {
			if levelofcolumns[idx] == rootLevel {
				continue
			}
			if idx < len(levelofcolumns)-1 {
				ctrl.next = trimArrayPattern(levelofcolumns[idx+1])
			} else {
				for key := range ctrl.value {
					ctrl.next = trimArrayPattern(key)
					break
				}
			}

			var script string //一個欄位的單一層級更新語法
			if isArray(levelofcolumns[idx]) {
				script = processArray(levelofcolumns[idx], arrayTemplete, ctrl)
			} else {
				script = processColumn(levelofcolumns[idx], columnTemplate, ctrl)
			}

			scripts = strings.ReplaceAll(scripts, "<merge>", script)
		}

		leaves := "" //處理最底層語法
		for leaf := range val {
			if isArray(leaf) {
				leaves += processArray(leaf, leafArrayTemplate, ctrl) + comma
			} else {
				leaves += processColumn(leaf, leafTemplate, ctrl) + comma
			}
		}

		scripts = strings.ReplaceAll(scripts, "<merge>", strings.Trim(leaves, comma))

		aql := strings.ReplaceAll(template, "<doc>", docItem)
		for k, v := range bindVars {
			aql = strings.ReplaceAll(aql, "@"+k, v.(string))
		}
		aql = strings.ReplaceAll(aql, "<expr>", scripts)

		_, err := hdr.db.Query(context.Background(), aql, nil)
		if err != nil {
			logger.Errorf("db/arangoHdr/NestedUpdate: %v", err)
			logger.Errorf("update aql:\n%v", aql)
			return errors.DatabaseUpdate
		}
		logger.Infof("version: %v, collection:%v finished", version, collection)
	}

	return 0
}

// NestedSelectSummary 依照傳入的欄位定義(json path), 查詢特定 collection 的指定欄位的總和，回傳 aql
func (hdr *arangoHdr) NestedSelectSummary(collection string, params []string) string {
	bindVars := map[string]interface{}{
		"@Collection": collection,
	}

	selectTemplate := `RETURN {
	<expr>
}`
	exprTemplate := `<leaf>: SUM(
	    FOR <doc> IN @@Collection
	<merge>
	)
	`
	arrayTemplate := `    FOR <item> IN <parent>.<aname>
	    <filter>
	<merge>`
	leafTemplate := `    RETURN <parent>`

	var expr string
	for idx := range params {
		levelofcolumns := getLevels(params[idx])
		ctrl := &parsectrl{
			parrent: docItem,
			alias:   make(map[string]string),
		}
		if len(levelofcolumns) == 0 {
			continue
		}

		scripts := strings.ReplaceAll(exprTemplate, "<doc>", docItem)
		for i, level := range levelofcolumns {
			aname := trimArrayPattern(level)
			if i == len(levelofcolumns)-1 {
				if isArray(level) {
					ctrl.idx++
					itemName := fmt.Sprint(arrayItemPrefix, ctrl.idx, arrayItemPostfix)
					script := strings.ReplaceAll(arrayTemplate, "<aname>", aname)
					filter := arrayFilter(level, ctrl)
					script = strings.ReplaceAll(script, "<filter>", filter)
					script = strings.ReplaceAll(script, "<parent>", ctrl.parrent)
					ctrl.parrent = itemName
					script = strings.ReplaceAll(script, "<item>", itemName)
					leafScript := strings.ReplaceAll(leafTemplate, "<parent>", itemName)
					script = strings.ReplaceAll(script, "<merge>", leafScript)
					scripts = strings.ReplaceAll(scripts, "<merge>", script)
					scripts = strings.ReplaceAll(scripts, "<leaf>", aname)
					continue
				}
				ctrl.parrent = fmt.Sprint(ctrl.parrent, string(dot), level)
				script := strings.ReplaceAll(leafTemplate, "<parent>", ctrl.parrent)
				scripts = strings.ReplaceAll(scripts, "<merge>", script)
				scripts = strings.ReplaceAll(scripts, "<leaf>", level)
				continue
			}

			if isArray(level) {
				ctrl.idx++
				aname := trimArrayPattern(level)
				itemName := fmt.Sprint(arrayItemPrefix, ctrl.idx, arrayItemPostfix)
				ctrl.alias[aname] = itemName
				script := strings.ReplaceAll(arrayTemplate, "<aname>", aname)
				script = strings.ReplaceAll(script, "<parent>", ctrl.parrent)
				ctrl.parrent = itemName
				filter := arrayFilter(level, ctrl)
				script = strings.ReplaceAll(script, "<filter>", filter)
				script = strings.ReplaceAll(script, "<item>", itemName)
				scripts = strings.ReplaceAll(scripts, "<merge>", script)
			} else {
				ctrl.parrent = fmt.Sprint(ctrl.parrent, string(dot), level)
			}
		}
		expr += scripts + comma
	}

	aql := strings.ReplaceAll(selectTemplate, "<expr>", strings.Trim(expr, comma))
	for k, v := range bindVars {
		aql = strings.ReplaceAll(aql, "@"+k, v.(string))
	}
	return aql
}

func groupValues(params map[string]string) map[string]map[string]string {
	res := make(map[string]map[string]string)
	for key, val := range params {
		groupName, valName := getGroup(key)
		vmap, ex := res[groupName]
		if !ex {
			vmap = make(map[string]string)
		}
		vmap[valName] = val
		res[groupName] = vmap
	}
	return res
}

func getGroup(key string) (string, string) {
	for i := len(key) - 1; i >= 0; i-- {
		if key[i] == dot {
			groupName := key[:i]
			valName := key[i+1:]
			return groupName, valName
		}
	}
	return rootLevel, key
}

func getLevels(input string) []string {
	res := []string{}
	in := false
	cur := 0
	for i, s := range input {
		switch s {
		case leftSquareBracket:
			in = true
		case rightSquareBracket:
			in = false
		case dot:
			if !in {
				res = append(res, input[cur:i])
				cur = i + 1
			}
		}
	}
	if cur < len(input) {
		res = append(res, input[cur:])
	}
	return res
}

func trimArrayPattern(arrayDefine string) string {
	idx := strings.IndexRune(arrayDefine, leftSquareBracket)
	if idx > 0 {
		return arrayDefine[:idx]
	}
	return arrayDefine
}

func arrayFilter(arrayDefine string, ctrl *parsectrl) string {
	idx := strings.IndexRune(arrayDefine, leftSquareBracket)
	res := ""
	if idx < 0 {
		return res
	}
	cond := arrayDefine[idx+1 : len(arrayDefine)-1]
	if len(cond) == 0 {
		return res
	}
	switch cond[0] {
	case questionMark:
		// input : ?(processData.eventCode == 11)
		expr := cond[2 : len(cond)-1]
		expr = strings.ReplaceAll(expr, arraySelfItem, ctrl.parrent)
		for key, name := range ctrl.alias {
			expr = strings.ReplaceAll(expr, key, name)
		}
		res = strings.ReplaceAll(arrayFilterTemplate, "<expr>", expr)
	}
	return res
}

func isArray(name string) bool {
	return strings.ContainsRune(name, leftSquareBracket)
}

type parsectrl struct {
	next    string
	parrent string
	idx     int
	value   map[string]string
	alias   map[string]string
}

func processArray(define, template string, ctrl *parsectrl) string {
	ctrl.idx++
	aname := trimArrayPattern(define)
	itemName := fmt.Sprint(arrayItemPrefix, ctrl.idx, arrayItemPostfix)
	ctrl.alias[trimArrayPattern(define)] = itemName
	script := strings.ReplaceAll(template, "<aname>", aname)
	script = strings.ReplaceAll(script, "<item>", itemName)
	script = strings.ReplaceAll(script, "<next>", ctrl.next)
	if _, ex := ctrl.value[define]; !ex {
		script = strings.ReplaceAll(script, "<parrent>", ctrl.parrent)
		ctrl.parrent = itemName
	} else {
		script = strings.ReplaceAll(script, "<parrent>", fmt.Sprint(ctrl.parrent, string(dot), aname))
	}
	filter := arrayFilter(define, ctrl)
	script = strings.ReplaceAll(script, "<filter>", filter)
	script = strings.ReplaceAll(script, "<leaf>", leafArrayItem)
	value := strings.ReplaceAll(ctrl.value[define], aname, leafArrayItem)
	script = strings.ReplaceAll(script, "<value>", value)
	return script
}

func processColumn(define, template string, ctrl *parsectrl) string {
	script := strings.ReplaceAll(template, "<cname>", define)
	value := strings.ReplaceAll(ctrl.value[define], define, fmt.Sprint(ctrl.parrent, string(dot), define))
	script = strings.ReplaceAll(script, "<value>", value)
	script = strings.ReplaceAll(script, "<next>", ctrl.next)
	if _, ex := ctrl.value[define]; !ex {
		ctrl.parrent = fmt.Sprint(ctrl.parrent, string(dot), define)
	}
	script = strings.ReplaceAll(script, "<parrent>", ctrl.parrent)
	return script
}
