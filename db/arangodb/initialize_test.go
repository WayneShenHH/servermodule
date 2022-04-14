package arangodb

import (
	"testing"
	"time"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/gracefulshutdown"
)

func Test_NewArango(t *testing.T) {
	gracefulshutdown.Start()
	initialize(&config.ArangoDBConfig{
		Addr:          "http://127.0.0.1:8529",
		Database:      "Database",
		Connlimit:     5,
		RetryCount:    5,
		RetryInterval: time.Millisecond * 300,
		HttpProtocol:  "1.1",
	})
	time.Sleep(3 * time.Second)
	gracefulshutdown.Shutdown()
}
