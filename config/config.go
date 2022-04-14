package config

import (
	"time"

	"github.com/spf13/viper"
)

// Setting 全系統設定吃這個
var Setting Config

// Config config model of env
type Config struct {
	Database  *DatabaseConfig
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
	Level       string
	Formatter   string
	ServiceCode int
}

// RedisConfig redis setting
type RedisConfig struct {
	Host         string
	Port         int
	Index        int
	MaxConns     int
	MaxIdleConns int
	Timeout      int  // 單次連線超時 (Millisecond)
	IdleTimeout  int  // 連線池內存活超時 (Millisecond)
	ReadTimeout  int  // 讀取超時，應小於等於 Timeout (Millisecond)
	WriteTimeout int  // 寫入超時，應小於等於 Timeout (Millisecond)
	JWTExpire    uint // JWT 在 Redis 中的有效期限，單位為 Minute
	Block        bool // 當達 MaxConn 後收到 Get() 請求時是否等待
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
