package arangodb

import (
	"context"

	driver "github.com/arangodb/go-driver"

	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/protocol"
	"github.com/WayneShenHH/servermodule/protocol/constant"
	"github.com/WayneShenHH/servermodule/protocol/dao"
	"github.com/WayneShenHH/servermodule/protocol/errorcode"
	"github.com/WayneShenHH/servermodule/util"
)

// seperate steps by collection, check data exist before each step
func (hdr *arangoHdr) Seed() protocol.ErrorCode {
	code := hdr.addCollection(constant.Collection_Players)
	if code > 0 {
		return code
	}

	code = hdr.documentPlayers()
	if code > 0 {
		return code
	}

	return 0
}

func (hdr *arangoHdr) addCollection(name string) protocol.ErrorCode {
	exist, err := hdr.db.CollectionExists(context.TODO(), name)
	if err != nil {
		logger.Errorf("db/arangoHdr/collectionPlayers/CollectionExists: %v", err)
		return errorcode.DatabaseUpdate
	}

	if exist {
		return errorcode.Success
	}

	_, err = hdr.db.CreateCollection(context.TODO(), name, nil)
	if err != nil {
		logger.Errorf("db/arangoHdr/collectionPlayers/CreateCollection: %v", err)
		return errorcode.DatabaseCollection
	}

	return errorcode.Success
}

func (hdr *arangoHdr) documentPlayers() protocol.ErrorCode {
	action := `function (Params) {
		const players = JSON.parse(Params[1]);
		const db = require('@arangodb').db;
		const col = db._collection(Params[0]);
		for (var i = 0; i < players.length; ++i) {
			if (col.exists(players[i])) {
				continue;
			}
			col.save(players[i]);
		}
		return 1;
	}`

	players := []dao.Player{
		{
			Key:  "1",
			Name: "Admin",
		},
	}

	jsonData, err := util.Marshal(players)
	if err != nil {
		logger.Errorf("db/arangoHdr/util.Marshal: %v", err)
		return errorcode.InputValueInvalid
	}

	options := &driver.TransactionOptions{
		MaxTransactionSize: 1000,
		WriteCollections:   []string{constant.Collection_Players},
		ReadCollections:    []string{constant.Collection_Players},
		Params:             []interface{}{constant.Collection_Players, string(jsonData)},
		WaitForSync:        false,
	}

	_, err = hdr.db.Transaction(context.TODO(), action, options)
	if err != nil {
		logger.Errorf("db/arangoHdr/UpdateDocument/Transaction: %v", err)
		return errorcode.DatabaseUpdate
	}

	return 0
}
