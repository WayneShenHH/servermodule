package config

import (
	"time"

	"github.com/spf13/viper"
)

// Setting 全系統設定吃這個
var Setting Config

// Config config model of env
type Config struct {
	Database  *ArangoDBConfig
	Redis     *RedisConfig
	NSQ       *NSQConfig
	BigQuery  *BigQueryConfig
	MQ        *MQConnect
	Logger    *LoggerConfig
	Worker    *WorkerConfig
	GSuite    *GSuiteConfig
	Stomp     *StompConfig
	Websocket *WebsocketConfig
	HTTP      *HTTPConfig
	Gin       *GinConfig
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

// ArangoDBConfig 資料庫連線設定
type ArangoDBConfig struct {
	Addr          string        `yaml:"addr,omitempty"`
	Database      string        `yaml:"database,omitempty"`
	Connlimit     int           `yaml:"connlimit,omitempty"`
	Username      string        `yaml:"username,omitempty"`
	Password      string        `yaml:"password,omitempty"`
	RetryCount    int           `yaml:"retryCount,omitempty"`
	RetryInterval time.Duration `yaml:"retryInterval,omitempty"`
	HttpProtocol  string        `yaml:"httpProtocol,omitempty"`
}

// LoggerConfig logger setting
type LoggerConfig struct {
	Level       string
	Formatter   string
	ServiceCode int
}

// RedisConfig redis setting
type RedisConfig struct {
	Addr          string        `yaml:"addr,omitempty"`
	Password      string        `yaml:"password,omitempty"`
	PoolSize      int           `yaml:"poolSize,omitempty"`
	RetryCount    int           `yaml:"retryCount,omitempty"`
	RetryInterval time.Duration `yaml:"retryInterval,omitempty"`
}

// NatsConfig setting
type NatsConfig struct {
	Addr              string        `yaml:"addr,omitempty"`
	Username          string        `yaml:"username,omitempty"`
	Password          string        `yaml:"password,omitempty"`
	ClusterID         string        `yaml:"clusterID,omitempty"`
	ReconnInterval    time.Duration `yaml:"reconnInterval,omitempty"`
	ConnectTimeOut    time.Duration `yaml:"connectTimeOut,omitempty"`
	StanPingsInterval int           `yaml:"stanPingsInterval,omitempty"`
	StanPingsMaxOut   int           `yaml:"stanPingsMaxOut,omitempty"`
}

// NSQConfig config
type NSQConfig struct {
	NSQDTCP       string //:4150 , producer
	NSQDHTTP      string //:4151
	NSQDValid     bool   // 是否在啟動的時候驗證 nsqd 的連線狀況
	NSQLookupTCP  string //:4160
	NSQLookupHTTP string //:4161 , consumer
	Concurrency   int    // consumer concurrency handler count
	MaxInFlight   int    // 最大可以同時連結的 nsqd 數量
}

// BigQueryConfig bigquery config
type BigQueryConfig struct {
	ProjectID string
	DatasetID string
}

// MQConnect MQ connection setting
type MQConnect struct {
	Host     string
	Port     int
	Username string
	Password string
}

// WorkerConfig worker alert by latency of offer
type WorkerConfig struct {
	LatencyWarning int // 設定訊息延遲超過多少毫秒（millisecond），記錄警告
}

// GSuiteConfig 串接線上 google gsuite 資料來源
type GSuiteConfig struct {
	Cert string
}

// StompConfig stomp
type StompConfig struct {
	SSL bool // Use ssl connection or not
}

// GRPCConfig config
type GRPCConfig struct {
	Addr string
}

// WebsocketConfig ws 設定
type WebsocketConfig struct {
	Addr         string
	PingDelay    int
	PingURL      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// HTTPConfig 的設定檔結構
type HTTPConfig struct {
	Addr      string
	PingURL   string
	BaseURL   string
	TimeoutMS uint
}

// GinConfig  GIN 套件相關設定
type GinConfig struct {
	LogEnable bool
	CacheMode string
}
