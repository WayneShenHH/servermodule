package logger

import (
	"github.com/WayneShenHH/servermodule/logger/env"

	"github.com/WayneShenHH/servermodule/logger/heavenlogger"
)

// Logger interface
type Logger interface {
	OpenFile(fileName string)
	SetFatalCallback(fn func(msg string))

	WarnCallStack(args ...interface{})
	InfoCallStack(msg ...interface{})
	DebugCallStack(msg ...interface{})

	Fatal(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(msg ...interface{})
	Debug(msg ...interface{})
	Debug4(msg ...interface{})
	Debug3(msg ...interface{})
	Debug2(msg ...interface{})
	Debug1(msg ...interface{})

	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Debug4f(format string, args ...interface{})
	Debug3f(format string, args ...interface{})
	Debug2f(format string, args ...interface{})
	Debug1f(format string, args ...interface{})

	ApmError(traceId, transactionId, spanId, msg string)
	ApmWarn(traceId, transactionId, spanId, msg string)
	ApmInfo(traceId, transactionId, spanId, msg string)
	ApmDebug(traceId, transactionId, spanId, msg string)
	ApmDebug1f(traceId, transactionId, spanId, msg string)
	ApmFatalf(traceId, transactionId, spanId, msg string)
}

var instance Logger

// Init logger by level & formatter & service-code
//
// default is info-level & json-format while parameter invalid
func Init(level, formatter string, code int) {
	instance = heavenlogger.New(level, formatter, code)
}

// OpenFile output to file
func OpenFile(fileName string) {
	instance.OpenFile(fileName)
}

func init() {
	Init(env.Setting.Level, env.Setting.Formatter, env.Setting.Code) // default mode
}

// SetFatalCallback config callback task, which run before fatal.panic
func SetFatalCallback(fn func(msg string)) {
	instance.SetFatalCallback(fn)
}

// WarnCallStack console warning & stack-trace
func WarnCallStack(args ...interface{}) {
	instance.WarnCallStack(args...)
}

// InfoCallStack console info & stack-trace
func InfoCallStack(args ...interface{}) {
	instance.InfoCallStack(args...)
}

// DebugCallStack console debug log & stack-trace
func DebugCallStack(args ...interface{}) {
	instance.DebugCallStack(args...)
}

// Fatal console & panic
func Fatal(args ...interface{}) {
	instance.Fatal(args...)
}

// Error console when err isn't null
func Error(args ...interface{}) {
	instance.Error(args...)
}

// Warn console warning
func Warn(args ...interface{}) {
	instance.Warn(args...)
}

// Info console info
func Info(args ...interface{}) {
	instance.Info(args...)
}

// Debug console debug log
func Debug(args ...interface{}) {
	instance.Debug(args...)
}

// Debug console debug log
func Debug4(args ...interface{}) {
	instance.Debug4(args...)
}

// Debug console debug log
func Debug3(args ...interface{}) {
	instance.Debug3(args...)
}

// Debug console debug log
func Debug2(args ...interface{}) {
	instance.Debug2(args...)
}

// Debug console debug log
func Debug1(args ...interface{}) {
	instance.Debug1(args...)
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

// Debugf equal to fmt.Printf but auto given log struct
func Debug4f(format string, args ...interface{}) {
	instance.Debug4f(format, args...)
}

// Debugf equal to fmt.Printf but auto given log struct
func Debug3f(format string, args ...interface{}) {
	instance.Debug3f(format, args...)
}

// Debugf equal to fmt.Printf but auto given log struct
func Debug2f(format string, args ...interface{}) {
	instance.Debug2f(format, args...)
}

// Debugf equal to fmt.Printf but auto given log struct
func Debug1f(format string, args ...interface{}) {
	instance.Debug1f(format, args...)
}

func ApmError(traceId, transactionId, spanId, msg string) {
	instance.ApmError(traceId, transactionId, spanId, msg)
}
func ApmWarn(traceId, transactionId, spanId, msg string) {
	instance.ApmWarn(traceId, transactionId, spanId, msg)
}
func ApmInfo(traceId, transactionId, spanId, msg string) {
	instance.ApmInfo(traceId, transactionId, spanId, msg)
}
func ApmDebug(traceId, transactionId, spanId, msg string) {
	instance.ApmDebug(traceId, transactionId, spanId, msg)
}

func ApmDebug1f(traceId, transactionId, spanId, msg string) {
	instance.ApmDebug1f(traceId, transactionId, spanId, msg)
}

func ApmFatalf(traceId, transactionId, spanId, msg string) {
	instance.ApmFatalf(traceId, transactionId, spanId, msg)
}
