package logger_test

import (
	"fmt"
	"testing"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger"
)

func Test_Info(t *testing.T) {
	logger.Init(&config.LoggerConfig{
		StdLevel:   config.Debug,
		LoggerName: config.Logrus,
		Formatter:  config.Stackdriver,
	})

	msgStr := "testing string"
	msgStruct := map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	err := fmt.Errorf("error content")

	logger.Debug()
	logger.Debug(msgStr, err, 503)
	logger.Debug(msgStr)
	logger.Debugf("some where got an error: %v", err)
	logger.Info(msgStr, err, msgStruct, 503)
	logger.Info(msgStr, 503)
	logger.Warn(msgStr, err, msgStruct, 503)
	logger.Error(msgStr, err, msgStruct, 503)
	logger.Error(503)
	logger.Error(msgStr)
	logger.Error(err)
	logger.Error(msgStruct, 503)
}
