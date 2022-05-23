package heavenlogger

import (
	"os"

	"github.com/WayneShenHH/servermodule/gracefulshutdown"
	"github.com/WayneShenHH/servermodule/slackalert"
)

// SetFatalCallback config fatal callback
func (l *Logger) SetFatalCallback(fn func(msg string)) {
	l.fatalCallback = fn
}

// OpenFile output to file
func (l *Logger) OpenFile(fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	l.outputFile = f
}

// WarnCallStack ...
func (l *Logger) WarnCallStack(args ...interface{}) {
	e := l.log(warnLevel, true, args...)
	if slackalert.IsSlackEnable() {
		slackalert.SendWarns(e)
	}
}

// InfoCallStack ...
func (l *Logger) InfoCallStack(args ...interface{}) {
	l.log(infoLevel, true, args...)
}

// DebugCallStack console with stack trace
func (l *Logger) DebugCallStack(args ...interface{}) {
	l.log(debugLevel, true, args...)
}

// Fatal ...
func (l *Logger) Fatal(args ...interface{}) {
	e := l.log(fatalLevel, true, args...)

	if slackalert.IsSlackEnable() {
		slackalert.SendError(e)
	}

	if l.fatalCallback != nil {
		l.fatalCallback(e)
	}

	gracefulshutdown.Shutdown()
}

// Error ...
func (l *Logger) Error(args ...interface{}) {
	e := l.log(errorLevel, true, args...)
	if slackalert.IsSlackEnable() {
		slackalert.SendWarns(e)
	}
}

// Warn ...
func (l *Logger) Warn(args ...interface{}) {
	e := l.log(warnLevel, false, args...)
	if slackalert.IsSlackEnable() {
		slackalert.SendWarns(e)
	}
}

// Info ...
func (l *Logger) Info(args ...interface{}) {
	l.log(infoLevel, false, args...)
}

// Debug ...
func (l *Logger) Debug(args ...interface{}) {
	l.log(debugLevel, false, args...)
}

// Debug ...
func (l *Logger) Debug4(args ...interface{}) {
	l.log(debug4Level, false, args...)
}

// Debug ...
func (l *Logger) Debug3(args ...interface{}) {
	l.log(debug3Level, false, args...)
}

// Debug ...
func (l *Logger) Debug2(args ...interface{}) {
	l.log(debug2Level, false, args...)
}

// Debug ...
func (l *Logger) Debug1(args ...interface{}) {
	l.log(debug1Level, false, args...)
}

// Fatalf equal to fmt.Printf but auto given log prefix
func (l *Logger) Fatalf(format string, args ...interface{}) {
	e := l.logf(fatalLevel, format, true, args...)
	if slackalert.IsSlackEnable() {
		slackalert.SendWarns(e)
	}

	if l.fatalCallback != nil {
		l.fatalCallback(e)
	}

	gracefulshutdown.Shutdown()
}

// Errorf equal to fmt.Printf but auto given log prefix
func (l *Logger) Errorf(format string, args ...interface{}) {
	e := l.logf(errorLevel, format, true, args...)
	if slackalert.IsSlackEnable() {
		slackalert.SendWarns(e)
	}
}

// Warnf equal to fmt.Printf but auto given log prefix
func (l *Logger) Warnf(format string, args ...interface{}) {
	e := l.logf(warnLevel, format, false, args...)
	if slackalert.IsSlackEnable() {
		slackalert.SendWarns(e)
	}
}

// Infof equal to fmt.Printf but auto given log prefix
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logf(infoLevel, format, false, args...)
}

// Debugf equal to fmt.Printf but auto given log prefix
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logf(debugLevel, format, false, args...)
}

// Debugf equal to fmt.Printf but auto given log prefix
func (l *Logger) Debug4f(format string, args ...interface{}) {
	l.logf(debug4Level, format, false, args...)
}

// Debugf equal to fmt.Printf but auto given log prefix
func (l *Logger) Debug3f(format string, args ...interface{}) {
	l.logf(debug3Level, format, false, args...)
}

// Debugf equal to fmt.Printf but auto given log prefix
func (l *Logger) Debug2f(format string, args ...interface{}) {
	l.logf(debug2Level, format, false, args...)
}

// Debugf equal to fmt.Printf but auto given log prefix
func (l *Logger) Debug1f(format string, args ...interface{}) {
	l.logf(debug1Level, format, false, args...)
}

// ApmError
func (l *Logger) ApmError(traceId, transactionId, spanId, msg string) {
	e := l.log(errorLevel, true, ApmTraceId(traceId), ApmTransactionId(transactionId), ApmSpanId(spanId), msg)
	if slackalert.IsSlackEnable() {
		slackalert.SendWarns(e)
	}
}

// ApmWarn
func (l *Logger) ApmWarn(traceId, transactionId, spanId, msg string) {
	e := l.log(warnLevel, false, ApmTraceId(traceId), ApmTransactionId(transactionId), ApmSpanId(spanId), msg)
	if slackalert.IsSlackEnable() {
		slackalert.SendWarns(e)
	}
}

// ApmInfo
func (l *Logger) ApmInfo(traceId, transactionId, spanId, msg string) {
	l.log(infoLevel, false, ApmTraceId(traceId), ApmTransactionId(transactionId), ApmSpanId(spanId), msg)
}

// ApmDebug
func (l *Logger) ApmDebug(traceId, transactionId, spanId, msg string) {
	l.log(debugLevel, false, ApmTraceId(traceId), ApmTransactionId(transactionId), ApmSpanId(spanId), msg)
}

// ApmDebug1f
func (l *Logger) ApmDebug1f(traceId, transactionId, spanId, msg string) {
	l.log(debug1Level, false, ApmTraceId(traceId), ApmTransactionId(transactionId), ApmSpanId(spanId), msg)
}

// ApmFatalf
func (l *Logger) ApmFatalf(traceId, transactionId, spanId, msg string) {
	e := l.log(fatalLevel, false, ApmTraceId(traceId), ApmTransactionId(transactionId), ApmSpanId(spanId), msg)
	if slackalert.IsSlackEnable() {
		slackalert.SendError(e)
	}

	if l.fatalCallback != nil {
		l.fatalCallback(e)
	}
	gracefulshutdown.Shutdown()
}
