package logger

import (
	"fmt"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/logger/logrushdr"
	"github.com/WayneShenHH/servermodule/logger/zaphdr"
)

// Logger interface
type Logger interface {
	Debug(fields ...interface{})
	Info(fields ...interface{})
	Warn(fields ...interface{})
	Error(fields ...interface{})
	Fatal(fields ...interface{})
}

var instance Logger

// Init logger by name
func Init(cfg *config.LoggerConfig) {
	switch cfg.LoggerName {
	case config.Zap:
		instance = zaphdr.New(cfg)
	case config.Logrus:
		instance = logrushdr.New(cfg)
	default:
		instance = logrushdr.New(cfg)
	}
}

// Debug log
func Debug(fields ...interface{}) {
	instance.Debug(fields...)
}

// Info log
func Info(fields ...interface{}) {
	instance.Info(fields...)
}

// Warn log
func Warn(fields ...interface{}) {
	instance.Warn(fields...)
}

// Error log
func Error(fields ...interface{}) {
	instance.Error(fields...)
}

// Fatal log and os.Exit(1)
func Fatal(fields ...interface{}) {
	instance.Fatal(fields...)
}

// Debugf log
func Debugf(format string, args ...interface{}) {
	instance.Debug(fmt.Sprintf(format, args...))
}

// Infof log
func Infof(format string, args ...interface{}) {
	instance.Info(fmt.Sprintf(format, args...))
}

// Warnf log
func Warnf(format string, args ...interface{}) {
	instance.Warn(fmt.Sprintf(format, args...))
}

// Errorf log
func Errorf(format string, args ...interface{}) {
	instance.Error(fmt.Sprintf(format, args...))
}

// Fatalf log and os.Exit(1)
func Fatalf(format string, args ...interface{}) {
	instance.Fatal(fmt.Sprintf(format, args...))
}
