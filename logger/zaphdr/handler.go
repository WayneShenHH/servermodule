package zaphdr

// Logger implemeent
type Logger struct{}

// Debug log Fatal
func (*Logger) Debug(msg ...interface{}) {
	Debug(msg)
}

// Info console log err
func (*Logger) Info(fields ...interface{}) {
	Info(fields...)
}

// Warn console log err
func (*Logger) Warn(fields ...interface{}) {
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
