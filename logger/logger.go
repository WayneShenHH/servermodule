package logger

import (
	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/constant"
	"github.com/WayneShenHH/servermodule/logger/zaphdr"
)

// logger mode list
const (
	ProdutionLevel   constant.Level = "info"
	DevelopmentLevel constant.Level = "debug"
	ErrorLevel       constant.Level = "error"

	JsonFormatter constant.LogFormatter = "json"
	StdFormatter  constant.LogFormatter = "std"
)

// Logger interface
type Logger interface {
	Fatal(args ...interface{})
	Error(args ...interface{})
	WarnCallStack(args ...interface{})
	Warn(args ...interface{})
	InfoCallStack(msg ...interface{})
	Info(msg ...interface{})
	Debug(msg ...interface{})
	DebugCallStack(msg ...interface{})
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	OpenFile(fileName string)

	SetFatalCallback(fn func(msg string))
	SetServiceCode(code constant.ServiceCode)
}

var instance Logger

// Init logger by level
func Init(cfg *config.LoggerConfig) {
	instance = zaphdr.New(cfg)
}

// OpenFile output to file
func OpenFile(fileName string) {
	instance.OpenFile(fileName)
}

func init() {
	Init(config.Setting.Logger) // default mode
}

// SetServiceCode constant code
func SetServiceCode(code constant.ServiceCode) {
	instance.SetServiceCode(code)
}

// SetFatalCallback constant callback task, which run before fatal.panic
func SetFatalCallback(fn func(msg string)) {
	instance.SetFatalCallback(fn)
}

// FatalOnFail console & panic when not success, auto filled service-code field when arg was constant.ServiceCode type
func Fatal(args ...interface{}) {
	instance.Fatal(args...)
}

// Error console when err isn't null, auto filled service-code field when arg was constant.ServiceCode type
func Error(args ...interface{}) {
	instance.Error(args...)
}

// WarnCallStack console warning & stack-trace, auto filled service-code field when arg was constant.ServiceCode type
func WarnCallStack(args ...interface{}) {
	instance.WarnCallStack(args...)
}

// Warn console warning, auto filled service-code field when arg was constant.ServiceCode type
func Warn(args ...interface{}) {
	instance.Warn(args...)
}

// InfoCallStack console info & stack-trace, auto filled service-code field when arg was constant.ServiceCode type
func InfoCallStack(args ...interface{}) {
	instance.InfoCallStack(args...)
}

// Info console info, auto filled service-code field when arg was constant.ServiceCode type
func Info(args ...interface{}) {
	instance.Info(args...)
}

// Debug console debug log, auto filled service-code field when arg was constant.ServiceCode type
func Debug(args ...interface{}) {
	instance.Debug(args...)
}

// DebugCallStack console debug log & stack-trace, auto filled service-code field when arg was constant.ServiceCode type
func DebugCallStack(args ...interface{}) {
	instance.DebugCallStack(args...)
}

// Fatalf equal to fmt.Printf but auto given log struct
func Fatalf(format string, args ...interface{}) {
	instance.Fatalf(format, args...)
}

// Errorf equal to fmt.Printf but auto given log struct
func Errorf(format string, args ...interface{}) {
	instance.Errorf(format, args...)
}

// Warnf equal to fmt.Printf but auto given log struct
func Warnf(format string, args ...interface{}) {
	instance.Warnf(format, args...)
}

// Infof equal to fmt.Printf but auto given log struct
func Infof(format string, args ...interface{}) {
	instance.Infof(format, args...)
}

// Debugf equal to fmt.Printf but auto given log struct
func Debugf(format string, args ...interface{}) {
	instance.Debugf(format, args...)
}
