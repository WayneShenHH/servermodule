package logger_test

import (
	"fmt"
	"testing"

	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/logger/heavenlogger"
	"github.com/WayneShenHH/servermodule/util/color"
)

type MsgStruct struct{}

var (
	msgStr    = "log-message-string"
	tagStr    = "a-tag"
	msgStruct = map[string]interface{}{
		"int": 1,
		"str": "string",
	}
	msgSlice          = []string{"item1", "item2"}
	err               = fmt.Errorf("error-message-string")
	serviceCode       = 500
	emptyMsgStructPtr *MsgStruct
)

func init() {
	logger.Init(heavenlogger.DebugLevel, heavenlogger.ConsoleFormatter, serviceCode)
	// logger.Init(heavenlogger.DebugLevel, heavenlogger.JsonFormatter)
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
	logger.Error(err, msgSlice, msgStruct, nil, emptyMsgStructPtr)
}
func Test_FormatError(t *testing.T) {
	logger.Errorf("error: %v", err)
	logger.Errorf("data: %v, %v", msgSlice, msgStruct)
	logger.Errorf("nil: %v", nil)
	logger.Errorf("emptyMsgStructPtr: %v", emptyMsgStructPtr)
}

func Test_Fatal(t *testing.T) {
	logger.Fatalf("error: %v", err)
}

func Test_OpenFile(t *testing.T) {
	logger.OpenFile("tmp.log")
	logger.Infof("Test_OpenFile Infof")
	logger.Errorf("Test_OpenFile Errorf")
}

func Test_Color(t *testing.T) {
	for i := 30; i < 107; i++ {
		for j := 30; j < 107; j++ {
			// if j != 44 {
			// 	continue
			// }
			if i >= 48 && i <= 89 {
				continue
			}
			if j >= 48 && j <= 89 {
				continue
			}
			c := color.New(color.Color(i), color.Color(j))
			msg := fmt.Sprintf("[%3d,%3d]", i, j)
			fmt.Print(c.Add(msg))
		}
	}
}

func Test_DebugLevel(t *testing.T) {
	logger.Debug1("hello")
	logger.Debug2("hello")
	logger.Debug3("hello")
	logger.Debug4("hello")
}
