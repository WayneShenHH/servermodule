package config

import (
	"time"

	"github.com/spf13/viper"
)

// Level of logger
type Level string

// LogFormatter of logger
type LogFormatter string

// Logger type name list
const (
	Zap    = "zap"
	Logrus = "logrus"
)

// LogFormatter list
const (
	Stackdriver LogFormatter = "stackdriver"
	File        LogFormatter = "file"
)

// Level list
const (
	Debug   Level = "debug"
	Info    Level = "info"
	Warning Level = "warning"
	Error   Level = "error"
	Fatal   Level = "fatal"
)

// Setting 全系統設定吃這個
var Setting Config

// Config config model of env
type Config struct {
	Database *DatabaseConfig
	Logger   *LoggerConfig
	GRPC     *GRPCConfig
}

// SetConfig get config by env setting
func SetConfig() {
	var c Config
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
	Setting = c
}

// DatabaseConfig 資料庫連線設定
type DatabaseConfig struct {
	Username        string
	Password        string
	Name            string
	Host            string
	MaxConns        int
	TimeoutDuration time.Duration
}

// LoggerConfig logger setting
type LoggerConfig struct {
	StdLevel   Level
	FileLevel  Level
	Formatter  LogFormatter
	LoggerName string
}

// GRPCConfig config
type GRPCConfig struct {
	Addr string
}
