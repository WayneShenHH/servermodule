package arangodb

import (
	"context"

	driver "github.com/arangodb/go-driver"

	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/protocol"
	"github.com/WayneShenHH/servermodule/protocol/errorcode"
	"github.com/WayneShenHH/servermodule/util"
)

// TODO: migrate logic
func (hdr *arangoHdr) Migrate(collection, key string, data interface{}) protocol.ErrorCode {
	jsonData, err := util.Marshal(data)
	if err != nil {
		logger.Debugf("db/arangoHdr/util.Marshal: %v", err)
		return errorcode.InputValueInvalid
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
		return errorcode.DatabaseUpdate
	}

	return 0
}
