// Package logger 系統共用 stdout logger
//nolint:unused // 先保留 logger 介面方法
package logger

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger/stackdriver"
	"github.com/WayneShenHH/servermodule/versions"
)

// std represents the logger which outputs to the stdout.
var std = logrus.New()
var workingDir string

// file represents the logger which outputs to the log file.
// var file = logrus.New()

// formatter formats the output format.
type formatter struct {
	isStdout bool
}

// Init initializes the global logger.
func Init(config *config.LoggerConfig) {
	// var fileFmt logrus.Formatter

	// Create the formatter for the both outputs.
	stdFmt := &formatter{true}
	// fileFmt = &formatter{false}

	// Std logger.
	std.Out = os.Stdout
	std.Level = logrus.DebugLevel
	std.Formatter = stdFmt
	// See https://github.com/sirupsen/logrus/issues/63#issuecomment-439946727
	// std.SetReportCaller(true)

	// prefixFormat := new(prefixed.TextFormatter)
	// prefixFormat.TimestampFormat = "2006-01-02 15:04:05"
	// prefixFormat.FullTimestamp = true
	// std.Formatter = prefixFormat

	// 設定輸出 format
	if strings.ToLower(config.Formatter) == "stackdriver" {
		std.Formatter = stackdriver.NewFormatter(
			stackdriver.WithService(versions.Name()),
			stackdriver.WithVersion(versions.Format()),
		)
	}

	// 設定 log 輸出等級
	switch strings.ToLower(config.StdLevel) {
	case "debug":
		std.Level = logrus.DebugLevel
	case "info":
		std.Level = logrus.InfoLevel
	case "warning":
		std.Level = logrus.WarnLevel
	case "error":
		std.Level = logrus.ErrorLevel
	case "fatal":
		std.Level = logrus.FatalLevel
	}

	if dir, err := os.Getwd(); err == nil {
		workingDir = dir
	}

	// File logger, create the log file when the file doesn't exist.
	// if _, err := os.Stat("./service.log"); os.IsNotExist(err) {
	// 	_, err := os.Create("./service.log")
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// // Open the log file so we can write the text to it.
	// f, err := os.OpenFile("./service.log", os.O_APPEND|os.O_WRONLY, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// file.Out = f
	// file.Level = logrus.FatalLevel
	// file.Formatter = fileFmt

	// Show debug message if it's debug mode.
	// if c.Bool("debug") {
	// 	std.Level = logrus.DebugLevel
	// }
}

// Format the input log.
func (f *formatter) Format(e *logrus.Entry) ([]byte, error) {
	// Implode the data to string with `k=v` format.
	dataString := ""
	if len(e.Data) != 0 {
		for k, v := range e.Data {
			dataString += fmt.Sprintf("%s=%+v ", k, v)
		}
		// Trim the trailing whitespace.
		dataString = dataString[0 : len(dataString)-1]
	}
	// Level like: DEBU, INFO, WARN, ERRO, FATA.
	level := strings.ToUpper(e.Level.String())[0:4]
	// Get the time with YYYY-mm-dd H:i:s format.
	time := e.Time.Format("2006-01-02 15:04:05")
	// Get the message.
	msg := e.Message
	// Get file name
	filename := ""
	// Get line number
	line := 0
	// funcname := ""
	// var pc uintptr

	// if runtimePc, runtimeFile, runtimeLine, ok := runtime.Caller(10); ok {
	if _, runtimeFile, runtimeLine, ok := runtime.Caller(10); ok {
		// runtimeFuncname := runtime.FuncForPC(runtimePc).Name()
		// fmt.Print(runtimePc, runtimeFile, runtimeLine, runtimeFuncname)
		// pc = runtimePc
		// funcname = runtimeFuncname
		line = runtimeLine
		filename = "." + strings.TrimPrefix(runtimeFile, workingDir)
	}

	// Set the color of the level with STDOUT.
	stdLevel := ""
	switch level {
	case "DEBU":
		stdLevel = color.New(color.FgBlue).Sprint(level)
	case "INFO":
		stdLevel = color.New(color.FgCyan).Sprint(level)
	case "WARN":
		stdLevel = color.New(color.FgYellow).Sprint(level)
	case "ERRO":
		stdLevel = color.New(color.FgRed).Sprint(level)
	case "FATA":
		stdLevel = color.New(color.FgHiRed).Sprint(level)
	}

	body := fmt.Sprintf("%s[%s] (%s:%d) %s ", level, time, filename, line, msg)
	data := dataString
	// Use the color level if it's STDOUT.
	if f.isStdout {
		body = fmt.Sprintf("%s[%s] (%s:%d) %s", stdLevel, time, filename, line, msg)
		// data = ""
	}
	// Hide the data if there's no data.
	if len(e.Data) == 0 {
		data = ""
	}

	// Mix the body and the data.
	output := fmt.Sprintf("%s %s\n", body, data)

	return []byte(output), nil
}

// DebugFields log debug
func DebugFields(msg string, fields logrus.Fields) {
	Fields(fields, "Debug", msg)
}

// InfoFields log info
func InfoFields(msg string, fields logrus.Fields) {
	Fields(fields, "Info", msg)
}

// WarningFields log warning
func WarningFields(msg string, fields logrus.Fields) {
	Fields(fields, "Warning", msg)
}

// ErrorFields log error
func ErrorFields(msg string, fields logrus.Fields) {
	Fields(fields, "Error", msg)
}

// FatalFields log fatel
func FatalFields(msg string, fields logrus.Fields) {
	Fields(fields, "Fatal", msg)
}

// Debug log debug
func Debug(msg interface{}) {
	Message("Debug", msg)
}

// Debugf log debug with format
func Debugf(format string, a ...interface{}) {
	Debug(fmt.Sprintf(format, a...))
}

// Info log info
func Info(msg interface{}) {
	Message("Info", msg)
}

// Infof log info with format
func Infof(format string, a ...interface{}) {
	Info(fmt.Sprintf(format, a...))
}

// Warning log warning
func Warning(msg interface{}) {
	Message("Warning", msg)
}

// Warningf log warning with format
func Warningf(format string, a ...interface{}) {
	Warning(fmt.Sprintf(format, a...))
}

// Error log error
func Error(msg interface{}) {
	Message("Error", msg)
}

// ErrorE log error and return
func ErrorE(err error) error {
	if err != nil {
		Message("Error", err)
	}
	return err
}

// Errorf log error with format
func Errorf(format string, a ...interface{}) {
	Error(fmt.Sprintf(format, a...))
}

// Fatal log Fatal
func Fatal(msg interface{}) {
	Message("Fatal", msg)
}

// Fatalf log Fatal with format
func Fatalf(format string, a ...interface{}) {
	Fatal(fmt.Sprintf(format, a...))
}

// WithField 加入 log 資訊
func WithField(fields logrus.Fields) *logrus.Entry {
	return std.WithFields(fields)
}

// Fields all log Fields
func Fields(fields logrus.Fields, lvl string, msg string) {
	s := std.WithFields(fields)
	// f := file.WithFields(fields)

	switch lvl {
	case "Debug":
		s.Debug(msg)
		// f.Debug(msg)
	case "Info":
		s.Info(msg)
		// f.Info(msg)
	case "Warning":
		s.Warning(msg)
		// f.Warning(msg)
	case "Error":
		s.Error(msg)
		// f.Error(msg)
	case "Fatal":
		s.Fatal(msg)
		// f.Fatal(msg)
	}
}

// Message log Message
func Message(lvl string, msg interface{}) {
	switch lvl {
	case "Debug":
		std.Debug(msg)
		// file.Debug(msg)
	case "Info":
		std.Info(msg)
		// file.Info(msg)
	case "Warning":
		std.Warning(msg)
		// file.Warning(msg)
	case "Error":
		std.Error(msg)
		// file.Error(msg)
	case "Fatal":
		std.Fatal(msg)
		// file.Fatal(msg)
	}
}
