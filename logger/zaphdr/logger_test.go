package zaphdr_test

import (
	"fmt"
	"testing"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger/zaphdr"
)

func Test_Info(t *testing.T) {
	zaphdr.New(&config.LoggerConfig{
		StdLevel: config.Debug,
	})
	msgStr := "string"
	msgStruct := map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	err := fmt.Errorf("error")
	zaphdr.Debug(err)
	zaphdr.Debug(msgStr)
	zaphdr.Info(msgStruct, 123)
	zaphdr.Info(err, 504)
	zaphdr.Warn(err, 504)
	zaphdr.Error(msgStr, err, msgStruct, 503)
	zaphdr.Fatal(msgStr, err, msgStruct, 503)
}
