package logrushdr

import (
	"github.com/sirupsen/logrus"
)

// Logger implemeent
type Logger struct{}

// Debug log Fatal
func (*Logger) Debug(msg ...interface{}) {
	DebugFields(processFields(msg))
}

// Info log Fatal
func (*Logger) Info(msg ...interface{}) {
	InfoFields(processFields(msg))
}

// Warn log Fatal
func (*Logger) Warn(msg ...interface{}) {
	WarningFields(processFields(msg))
}

// Error log Fatal
func (*Logger) Error(msg ...interface{}) {
	ErrorFields(processFields(msg))
}

// Fatal log Fatal
func (*Logger) Fatal(msg ...interface{}) {
	FatalFields(processFields(msg))
}

func processFields(fields []interface{}) (string, logrus.Fields) {
	msgField := make(logrus.Fields)
	msg := []interface{}{}
	var res string
	for idx := range fields {
		switch val := fields[idx].(type) {
		case int:
			msgField["code"] = val
		case error:
			msgField["error"] = val.Error()
		case string:
			res = val
		default:
			msg = append(msg, fields[idx])
		}
	}

	if len(msg) > 0 {
		if len(msg) == 1 {
			msgField["obj"] = msg[0]
		} else {
			msgField["obj"] = msg
		}
	}

	return res, msgField
}
