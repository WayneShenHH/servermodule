package logger_test

import (
	"fmt"
	"testing"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/constant"
	"github.com/WayneShenHH/servermodule/logger"
)

type MsgStruct struct{}

var (
	msgStr    = "log-message-string"
	tagStr    = "a-tag"
	msgStruct = map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	msgSlice                               = []string{"item1", "item2"}
	err                                    = fmt.Errorf("error-message-string")
	serviceCode       constant.ServiceCode = 500
	emptyMsgStructPtr *MsgStruct
	cfg               *config.LoggerConfig = &config.LoggerConfig{
		StdLevel:    constant.Debug,
		Formatter:   constant.JsonFormatter,
		ServiceCode: serviceCode,
	}
)

func init() {
	logger.Init(cfg)
}

func Test_Debug(t *testing.T) {
	logger.Debug(msgStr, msgStruct, msgSlice, err, emptyMsgStructPtr)
	logger.Debug(msgStr, msgStruct, msgSlice, err, emptyMsgStructPtr)
	logger.Debugf("%v", msgStr)
	logger.DebugCallStack(msgStr, msgStruct)
}

func Test_Info(t *testing.T) {
	logger.Info(tagStr, msgStr, err, msgStruct)
	logger.Infof("[%v] %v %v %v", tagStr, msgStr, err, msgStruct)
	logger.InfoCallStack(tagStr, msgStr)
}

func Test_Warn(t *testing.T) {
	logger.Warn(msgStr, msgStruct)
	logger.Warnf("%v, %v", msgStr, msgStruct)
	logger.WarnCallStack(msgStr, msgStruct)
}

func Test_Error(t *testing.T) {
	logger.Error(msgStr, msgStruct)
}
func Test_FormatError(t *testing.T) {
	logger.Errorf("error: %v", err)
	logger.Errorf("data: %v, %v", msgSlice, msgStruct)
	logger.Errorf("nil: %v", nil)
	logger.Errorf("emptyMsgStructPtr: %v", emptyMsgStructPtr)
}

func Test_Fatal(t *testing.T) {
	logger.Fatal(err, msgStr, msgStruct) // exit here
}

func Test_OpenFile(t *testing.T) {
	logger.OpenFile("tmp.log")
	logger.Infof("Test_OpenFile Infof")
	logger.Errorf("Test_OpenFile Errorf")
}
