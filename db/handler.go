package db

import (
	"context"
	"encoding/json"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/constant/errors"
	"github.com/WayneShenHH/servermodule/logger"
)

type arangoHdr struct {
	db driver.Database
}

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
