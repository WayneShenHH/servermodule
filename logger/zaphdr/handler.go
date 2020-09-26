package zaphdr

import (
	"go.uber.org/zap"

	"github.com/WayneShenHH/servermodule/util/stack"
)

// Logger implemeent
type Logger struct{}

// Debug log Fatal
func (*Logger) Debug(fields ...interface{}) {
	Debug(fields...)
}

// Info console log err
func (*Logger) Info(fields ...interface{}) {
	Info(fields...)
}

// Warn console log err
func (*Logger) Warn(fields ...interface{}) {
	Warn(fields...)
}

// Warn console log err
func (*Logger) WarnStack(fields ...interface{}) {
	fields = append(fields, zap.String("stacktrace", stack.TakeStacktrace(callerSkipOffset-1)))
	Warn(fields...)
}

// Error console log err
func (*Logger) Error(fields ...interface{}) {
	Error(fields...)
}

// Fatal console log err
func (*Logger) Fatal(fields ...interface{}) {
	Fatal(fields...)
}
