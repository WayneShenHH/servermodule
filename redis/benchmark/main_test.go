package benchmark

import (
	"fmt"
	"testing"
	"time"

	"github.com/WayneShenHH/servermodule/config"
	"github.com/WayneShenHH/servermodule/gracefulshutdown"
	"github.com/WayneShenHH/servermodule/redis"
)

func initRedis() {
	gracefulshutdown.Start()
	redis.Initialize(&config.RedisConfig{
		Addr: "localhost:7001,localhost:7002,localhost:7003,localhost:7004,localhost:7005,localhost:7006",
	})
}

func BenchmarkRedisSet(b *testing.B) {
	now := time.Now()
	initRedis()
	b.SetParallelism(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			r, err := redis.GetRedis().Eval(setScript(), []string{"benchmark/set"}, 1000, 1, 2)
			fmt.Println(r, err)
		}
	})
	b.ReportAllocs()
	if b.N > 1 {
		b.Logf("Elasped:%v", time.Since(now))
	}
}
func BenchmarkRedisGet(b *testing.B) {
	now := time.Now()
	initRedis()

	b.SetParallelism(1)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			redis.GetRedis().Eval(getScript(), []string{"benchmark/get"}, 500)
		}
	})
	b.ReportAllocs()
	if b.N > 1 {
		b.Logf("Elasped:%v", time.Since(now))
	}
}
