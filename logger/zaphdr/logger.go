package zaphdr

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/WayneShenHH/servermodule/config"
)

const (
	callerSkipOffset = 3 // zaphdr.Fatal & zaphdr.Logger.Fatal & logger.Fatal
	stackstraceKey   = "stacktrace"
)

var (
	instance *zap.Logger
	level    config.Level
)

// New initializes the global logger.
func New(cfg *config.LoggerConfig) *Logger {
	switch cfg.StdLevel {
	case config.Debug:
		setDev()
	case config.Info:
		setProd()
	}
	level = cfg.StdLevel
	return &Logger{}
}

func init() {
	setProd() //default prod mode
}

func setProd() {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      zapcore.OmitKey,
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     zapcore.OmitKey,
			StacktraceKey:  stackstraceKey,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochMillisTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
	instance, _ = cfg.Build(zap.AddCallerSkip(callerSkipOffset))
}

func setDev() {
	cfg := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "M",
			StacktraceKey:  stackstraceKey,
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.RFC3339TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}

	dev, _ := cfg.Build(zap.AddCallerSkip(callerSkipOffset))
	instance = dev
}

// Fatal console log err
func Fatal(fields ...interface{}) {
	message, fs := processFields(fields)
	instance.Fatal(message, fs...)
}

// Error console log err
func Error(fields ...interface{}) {
	message, fs := processFields(fields)
	instance.Error(message, fs...)
}

// Warn console log warn
func Warn(fields ...interface{}) {
	message, fs := processFields(fields)
	instance.Warn(message, fs...)
}

// Info console log info
func Info(fields ...interface{}) {
	message, fs := processFields(fields)
	instance.Info(message, fs...)
}

// Debug console log debug
func Debug(fields ...interface{}) {
	message, fs := processFields(fields)
	instance.Debug(message, fs...)
}

func processFields(fields []interface{}) (string, []zap.Field) {
	codekey := "serviceCode"
	msgKey := "msg"
	var msg []interface{}
	var msgField zap.Field
	var msgstr string
	res := []zap.Field{}
	for idx := range fields {
		switch val := fields[idx].(type) {
		case int:
			res = append(res, zap.Int(codekey, val))
		case error:
			msg = append(msg, val.Error())
		case string:
			msgstr = val
			if level != config.Debug {
				msg = append(msg, val)
			}
		case zap.Field:
			res = append(res, val)
		default:
			msg = append(msg, fields[idx])
		}
	}
	if len(msg) == 1 {
		msgField = zap.Any(msgKey, msg[0])
	} else {
		msgField = zap.Any(msgKey, msg)
	}
	if len(msg) > 0 {
		res = append(res, msgField)
	}
	return msgstr, res
}
