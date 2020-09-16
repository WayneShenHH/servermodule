package config

import "github.com/spf13/viper"

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
	Username string
	Password string
	Name     string
	Host     string
	MaxConns int
}

// LoggerConfig logger setting
type LoggerConfig struct {
	StdLevel  string
	FileLevel string
	Formatter string
}

// GRPCConfig config
type GRPCConfig struct {
	Addr string
}
