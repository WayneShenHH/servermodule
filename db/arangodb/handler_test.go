package arangodb

import (
	"testing"
	"time"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/gracefulshutdown"
	"github.com/WayneShenHH/servermodule/logger"
)

const info = `
{
	"firstName": "John",
	"lastName" : "doe",
	"age"      : 26,
	"address"  : {
	  "streetAddress": "naist street",
	  "city"         : "Nara",
	  "postalCode"   : "630-0192"
	},
	"phoneNumbers": [
	  {
		"type"  : "iPhone",
		"number": "0123-4567-8888"
	  },
	  {
		"type"  : "home",
		"number": "0123-4567-8910"
	  }
	],
	"banks": [
		{
			"name":"IBank",
			"balance":"100"
		},
		{
			"name":"CityBank",
			"balance":"890"
		}
	]
  }`

func init() {
	logger.Init("debug", "console", 100)
	gracefulshutdown.Start()
	initialize(&config.ArangoDBConfig{
		Addr:          "http://127.0.0.1:8529",
		Database:      "Database",
		Connlimit:     5,
		RetryCount:    5,
		RetryInterval: time.Millisecond * 300,
		HttpProtocol:  "1.1",
	})
}
func Test_NestedUpdate(t *testing.T) {
	db := GetArango()
	db.NestedUpdate("info", "1.01", map[string]string{
		"age":                               "age + 1",
		"banks[?(@.balance > 200)].balance": "balance * 100",
		// "banks[?(banks.balance > 200)].balance": "balance * 100",
		// "banks[*].balance": "balance * 100",
	})
}
