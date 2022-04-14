package redis

import (
	"testing"
	"time"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/gracefulshutdown"
	"github.com/WayneShenHH/servermodule/logger"
)

func Test_NewRedis(t *testing.T) {
	logger.Init("debug", "console", 0)
	gracefulshutdown.Start()
	defer gracefulshutdown.Shutdown()

	config := &config.RedisConfig{
		Addr:     "localhost:6379",
		Password: "",
		PoolSize: 20,
	}
	Initialize(config)
	GetRedis().SetGlobalUnLock()
	GetRedis().SetEX("test_key", "msg", time.Second)
	v, e := GetRedis().Get("test_key")
	logger.Debug("got:", v, e)
}
