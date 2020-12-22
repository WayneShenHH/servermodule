package zaphdr

import (
	"fmt"
	"os"

	"github.com/WayneShenHH/servermodule/constant"
)

// SetFatalCallback config fatal callback
func (l *Logger) SetFatalCallback(fn func(msg string)) {
	l.fatalCallback = fn
}

// SetServiceCode config code
func (l *Logger) SetServiceCode(code constant.ServiceCode) {
	l.serviceCode = code
}

// OpenFile output to file
func (l *Logger) OpenFile(fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	l.outputFile = f
}

// FatalOnError ...
func (l *Logger) Fatal(args ...interface{}) {
	l.fatal(args...)
}

// Error ...
func (l *Logger) Error(args ...interface{}) {
	l.error(args...)
}

// WarnCallStack ...
func (l *Logger) WarnCallStack(args ...interface{}) {
	l.warn(true, args...)
}

// Warn ...
func (l *Logger) Warn(args ...interface{}) {
	l.warn(false, args...)
}

// InfoCallStack ...
func (l *Logger) InfoCallStack(args ...interface{}) {
	l.info(true, args...)
}

// Info ...
func (l *Logger) Info(args ...interface{}) {
	l.info(false, args...)
}

// Debug ...
func (l *Logger) Debug(args ...interface{}) {
	l.debug(false, args...)
}

// DebugCallStack console with stack trace
func (l *Logger) DebugCallStack(args ...interface{}) {
	l.debug(true, args...)
}

// Fatalf equal to fmt.Printf but auto given log prefix
func (l *Logger) Fatalf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.fatal(msg)
}

// Errorf equal to fmt.Printf but auto given log prefix
func (l *Logger) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.error(msg)
}

// Warnf equal to fmt.Printf but auto given log prefix
func (l *Logger) Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.warn(false, msg)
}

// Infof equal to fmt.Printf but auto given log prefix
func (l *Logger) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.info(false, msg)
}

// Debugf equal to fmt.Printf but auto given log prefix
func (l *Logger) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	l.debug(false, msg)
}
