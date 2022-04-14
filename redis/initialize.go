package redis

import (
	"bytes"
	"context"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/gracefulshutdown"
	"github.com/WayneShenHH/servermodule/logger"
	"github.com/WayneShenHH/servermodule/protocol/constant"
	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"github.com/go-redis/redis/v8"
	"go.uber.org/atomic"
)

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

var handlerInstance *RedisHandler
var once sync.Once

// 專門處理 Redis 的相關事務
type RedisHandler struct {
	redisClient  redis.Cmdable
	ctx          context.Context
	config       *config.RedisConfig
	cacheManager *cache.Cache
	sync.Mutex
}

func switchConnect(redisIPs []string, poolSize int, config *config.RedisConfig) redis.Cmdable {
	switch len(redisIPs) {
	case 0: // error
		return nil
	case 1: // single
		return redis.NewClient(&redis.Options{
			Addr:            redisIPs[0],
			Password:        config.Password,
			PoolSize:        poolSize,
			MaxConnAge:      1 * time.Hour,
			MaxRetries:      config.RetryCount,
			MinRetryBackoff: config.RetryInterval,
			MaxRetryBackoff: config.RetryInterval,
		})
	default: // cluster
		return redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:      redisIPs,
			Password:   config.Password,
			PoolSize:   poolSize,
			MaxConnAge: 1 * time.Hour,
			// MaxRetries:      env.Setting.Redis.RetryCount,
			// MinRetryBackoff: env.Setting.Redis.RetryInterval,
			// MaxRetryBackoff: env.Setting.Redis.RetryInterval,
		})
	}
}

func retryLoop(ctx context.Context, shutdownChan, reTry chan struct{}, reTryCount *atomic.Uint32, handler *RedisHandler, config *config.RedisConfig) {
	for {
		select {
		case <-ctx.Done():
			// 斷掉 redis。
			shutdownChan <- struct{}{}
			return
		case <-reTry:
		ReTryFlag:
			if reTryCount.Inc() > uint32(config.RetryCount) {
				logger.Fatalf("Redis max retry count: %d", config.RetryCount)
				// 等待 shutdown 指令過來，然後等待關機。
				<-ctx.Done()
				shutdownChan <- struct{}{}
				return
			}
			redisClient, err := connect(ctx, config, reTry, shutdownChan)
			// 連線完成。
			if err == nil {
				reTryCount.Store(0)
				handler.redisClient = redisClient
			} else {
				logger.Warnf("Redis retry warn: %v", err.Error())
				goto ReTryFlag
			}
		}
	}
}

func GetRedis() *RedisHandler {
	if handlerInstance == nil {
		logger.Fatalf("GetRedis instance is nil, have to initialize redis")
		return nil
	}
	return handlerInstance
}

func Initialize(config *config.RedisConfig) {
	once.Do(func() {
		// 新增記憶體 cache
		ristrettoCache, err := ristretto.NewCache(&ristretto.Config{
			NumCounters: 1e6,     // 一百萬筆 cache if you expect your cache to hold 1,000,000 items when full, NumCounters should be 10,000,000 (10x)
			MaxCost:     1 << 29, // 536870912 五百MB if you want the cache to have a max capacity of 100MB, you would set MaxCost to 100,000,000
			BufferItems: 64,      // Unless you have a rare use case, using `64` as the BufferItems value results in good performance.
		})
		if err != nil {
			panic(err)
		}
		ristrettoStore := store.NewRistretto(ristrettoCache, nil)

		cacheManager := cache.New(ristrettoStore)
		ctx, shutdownChan := gracefulshutdown.GetContext(constant.Redis_Level)
		h := &RedisHandler{
			redisClient:  nil,
			ctx:          ctx,
			config:       config,
			cacheManager: cacheManager,
		}
		reTryCount := atomic.NewUint32(0)
		reTry := make(chan struct{}, 1)
	ReTryFlag:
		redisClient, err := connect(ctx, config, reTry, shutdownChan)
		if err == nil {
			h.redisClient = redisClient
		} else {
			logger.Warnf("Redis retry warn: %v", err.Error())
			if reTryCount.Inc() > uint32(config.RetryCount) {
				logger.Fatalf("Redis max retry count: %d", config.RetryCount)
				// 等待 shutdown 指令過來，然後等待關機。
				<-ctx.Done()
				shutdownChan <- struct{}{}
				return
			}
			goto ReTryFlag
		}
		go retryLoop(ctx, shutdownChan, reTry, reTryCount, h, config)
		logger.Info("Redis Initialize Done")
		handlerInstance = h
	})
}

func connect(ctx context.Context, config *config.RedisConfig, reTryChan, shutdownChan chan struct{}) (redis.Cmdable, error) {
	redisIPs := strings.Split(config.Addr, ",")
	for _, u := range redisIPs {
		_, err := url.Parse("http://" + u)
		if err != nil {
			logger.Fatalf("Redis IP url.Parse error: %v", err)
			// 等待 shutdown 指令過來，然後等待關機。
			<-ctx.Done()
			shutdownChan <- struct{}{}
			return nil, err
		}
	}

	client := switchConnect(redisIPs, config.PoolSize, config)
	if _, err := client.Ping(ctx).Result(); err != nil {
		logger.Warnf("redis Ping filed, err: %v", err.Error())
		return nil, err
	}

	// 每五秒 redis Ping 一次檢查。
	go pingLoop(ctx, client, reTryChan, shutdownChan)
	return client, nil
}

func pingLoop(ctx context.Context, redisClient redis.Cmdable, reTryChan, shutdownChan chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if _, err := redisClient.Ping(ctx).Result(); err != nil {
				logger.Warnf("redis Ping filed, err: %v", err.Error())
				// 送出 retry 訊號，跳出當下的 goroutine。
				reTryChan <- struct{}{}
				return
			}
		case <-ctx.Done():
			// 等待 shutdown 指令過來，然後等待關機。
			shutdownChan <- struct{}{}
			return
		}
	}
}
