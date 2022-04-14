package redis

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	Error_Redis_Handler_ZPopMin_Too_Many_Arguments = errors.New("Redis handler ZPopMin too many arguments")
	Error_Redis_Handler_ZPopMax_Too_Many_Arguments = errors.New("redis handler ZPopMax too many arguments")
	Error_Redis_Handler_BitPos_Too_Many_Arguments  = errors.New("redis handler bitPos too many arguments")
	Error_Redis_Handler_MemoryUsage                = errors.New("redis handler memoryUsage expects single sample count")
	Error_Redis_Null                               = errors.New("redis nil")
	Error_Redis_CanNotUsed                         = errors.New("redis can't used")
)

const canuseKey = "sport-can-use-fsizppqnd8u20dl-redis"

// 是否為 key 不存在的錯誤
func IsNullKeyError(err error) bool {
	return redis.Nil == err
}

func (h *RedisHandler) SetNoLock(key string, value interface{}, expiration time.Duration) (string, error) {
	return h.redisClient.Set(h.ctx, key, value, expiration).Result()
}

func (h *RedisHandler) Del(keys ...string) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.Del(h.ctx, keys...).Result()
}

func (h *RedisHandler) Get(key string) (string, error) {
	if h.CheckGlobalLock() {
		return "", Error_Redis_CanNotUsed
	}
	return h.redisClient.Get(h.ctx, key).Result()
}

func (h *RedisHandler) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	if h.CheckGlobalLock() {
		return "", Error_Redis_CanNotUsed
	}
	return h.redisClient.Set(h.ctx, key, value, expiration).Result()
}

func (h *RedisHandler) SetEX(key string, value interface{}, expiration time.Duration) (string, error) {
	if h.CheckGlobalLock() {
		return "", Error_Redis_CanNotUsed
	}
	return h.redisClient.SetEX(h.ctx, key, value, expiration).Result()
}

func (h *RedisHandler) HDel(key string, fields ...string) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.HDel(h.ctx, key, fields...).Result()
}

func (h *RedisHandler) HGet(key, field string) (string, error) {
	if h.CheckGlobalLock() {
		return "", Error_Redis_CanNotUsed
	}
	return h.redisClient.HGet(h.ctx, key, field).Result()
}

func (h *RedisHandler) HGetAll(key string) (map[string]string, error) {
	if h.CheckGlobalLock() {
		return nil, Error_Redis_CanNotUsed
	}
	return h.redisClient.HGetAll(h.ctx, key).Result()
}

func (h *RedisHandler) HKeys(key string) ([]string, error) {
	if h.CheckGlobalLock() {
		return []string{}, Error_Redis_CanNotUsed
	}
	return h.redisClient.HKeys(h.ctx, key).Result()
}

//
func (h *RedisHandler) HSet(key string, values ...interface{}) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.HSet(h.ctx, key, values...).Result()
}

func (h *RedisHandler) Pipeline() redis.Pipeliner {
	return h.redisClient.Pipeline()
}

func (h *RedisHandler) Expire(key string, expiration time.Duration) (bool, error) {
	if h.CheckGlobalLock() {
		return false, Error_Redis_CanNotUsed
	}
	return h.redisClient.Expire(h.ctx, key, expiration).Result()
}

// !!Redis 4.0 之後被棄用，請使用 HSet
// func (h *RedisHandler) HMSet(key string, values ...interface{}) (bool, error) {
// 	return h.redisClient.HMSet(h.ctx, key, values...).Result()
// }

func (h *RedisHandler) LRange(key string, start, stop int64) ([]string, error) {
	if h.CheckGlobalLock() {
		return []string{}, Error_Redis_CanNotUsed
	}
	return h.redisClient.LRange(h.ctx, key, start, stop).Result()
}

func (h *RedisHandler) Eval(script string, keys []string, args ...interface{}) (interface{}, error) {
	if h.CheckGlobalLock() {
		return nil, Error_Redis_CanNotUsed
	}
	return h.redisClient.Eval(h.ctx, script, keys, args...).Result()
}

func (h *RedisHandler) LPop(key string) (string, error) {
	if h.CheckGlobalLock() {
		return "", Error_Redis_CanNotUsed
	}
	return h.redisClient.LPop(h.ctx, key).Result()
}

func (h *RedisHandler) LIndex(key string, index int64) (string, error) {
	if h.CheckGlobalLock() {
		return "", Error_Redis_CanNotUsed
	}
	return h.redisClient.LIndex(h.ctx, key, index).Result()
}

func (h *RedisHandler) LPopCount(key string, count int) ([]string, error) {
	if h.CheckGlobalLock() {
		return []string{}, Error_Redis_CanNotUsed
	}
	return h.redisClient.LPopCount(h.ctx, key, count).Result()
}

func (h *RedisHandler) LPush(key string, values ...interface{}) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.LPush(h.ctx, key, values...).Result()
}

func (h *RedisHandler) LPushX(key string, values ...interface{}) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.LPushX(h.ctx, key, values...).Result()
}

func (h *RedisHandler) LRem(key string, count int64, value interface{}) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.LRem(h.ctx, key, count, value).Result()
}

func (h *RedisHandler) RPopLPush(source string, destination string) (string, error) {
	if h.CheckGlobalLock() {
		return "", Error_Redis_CanNotUsed
	}
	return h.redisClient.RPopLPush(h.ctx, source, destination).Result()
}

func (h *RedisHandler) BLPop(timeout time.Duration, keys ...string) ([]string, error) {
	if h.CheckGlobalLock() {
		return []string{}, Error_Redis_CanNotUsed
	}
	return h.redisClient.BLPop(h.ctx, timeout, keys...).Result()
}

func (h *RedisHandler) BRPop(timeout time.Duration, keys ...string) ([]string, error) {
	if h.CheckGlobalLock() {
		return []string{}, Error_Redis_CanNotUsed
	}
	return h.redisClient.BRPop(h.ctx, timeout, keys...).Result()
}

func (h *RedisHandler) BRPopLPush(source string, destination string, timeout time.Duration) (string, error) {
	if h.CheckGlobalLock() {
		return "", Error_Redis_CanNotUsed
	}
	return h.redisClient.BRPopLPush(h.ctx, source, destination, timeout).Result()
}

func (h *RedisHandler) HExists(key string, field string) (bool, error) {
	if h.CheckGlobalLock() {
		return false, Error_Redis_CanNotUsed
	}
	return h.redisClient.HExists(h.ctx, key, field).Result()
}

func (h *RedisHandler) IncrBy(key string, incr int64) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.IncrBy(h.ctx, key, incr).Result()
}

func (h *RedisHandler) HIncrBy(key string, field string, incr int64) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.HIncrBy(h.ctx, key, field, incr).Result()
}

func (h *RedisHandler) HIncrByFloat(key string, field string, incr float64) (float64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.HIncrByFloat(h.ctx, key, field, incr).Result()
}

func (h *RedisHandler) HSetNX(key string, field string, value interface{}) (bool, error) {
	if h.CheckGlobalLock() {
		return false, Error_Redis_CanNotUsed
	}
	return h.redisClient.HSetNX(h.ctx, key, field, value).Result()
}

func (h *RedisHandler) SetXX(key string, value interface{}, expiration time.Duration) (bool, error) {
	if h.CheckGlobalLock() {
		return false, Error_Redis_CanNotUsed
	}
	return h.redisClient.SetXX(h.ctx, key, value, expiration).Result()
}

func (h *RedisHandler) SAdd(key string, members ...interface{}) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.SAdd(h.ctx, key, members...).Result()
}

func (h *RedisHandler) SAddNoLock(key string, members ...interface{}) (int64, error) {
	return h.redisClient.SAdd(h.ctx, key, members...).Result()
}

func (h *RedisHandler) SCard(key string) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.SCard(h.ctx, key).Result()
}

func (h *RedisHandler) SIsMember(key string, member interface{}) (bool, error) {
	if h.CheckGlobalLock() {
		return false, Error_Redis_CanNotUsed
	}
	return h.redisClient.SIsMember(h.ctx, key, member).Result()
}

func (h *RedisHandler) SMembers(key string) ([]string, error) {
	if h.CheckGlobalLock() {
		return nil, Error_Redis_CanNotUsed
	}
	return h.redisClient.SMembers(h.ctx, key).Result()
}

func (h *RedisHandler) SPop(key string) (string, error) {
	if h.CheckGlobalLock() {
		return "", Error_Redis_CanNotUsed
	}
	return h.redisClient.SPop(h.ctx, key).Result()
}

func (h *RedisHandler) SPopN(key string, count int64) ([]string, error) {
	if h.CheckGlobalLock() {
		return nil, Error_Redis_CanNotUsed
	}
	return h.redisClient.SPopN(h.ctx, key, count).Result()
}

func (h *RedisHandler) SRandMember(key string) (string, error) {
	if h.CheckGlobalLock() {
		return "", Error_Redis_CanNotUsed
	}
	return h.redisClient.SRandMember(h.ctx, key).Result()
}

func (h *RedisHandler) SRandMemberN(key string, count int64) ([]string, error) {
	if h.CheckGlobalLock() {
		return nil, Error_Redis_CanNotUsed
	}
	return h.redisClient.SRandMemberN(h.ctx, key, count).Result()
}

func (h *RedisHandler) SRem(key string, members ...interface{}) (int64, error) {
	if h.CheckGlobalLock() {
		return 0, Error_Redis_CanNotUsed
	}
	return h.redisClient.SRem(h.ctx, key, members...).Result()
}

func (h *RedisHandler) FlushAll() (string, error) {
	return h.redisClient.FlushAll(h.ctx).Result()
}

func (h *RedisHandler) CheckGlobalLock() bool {
	_, err := h.redisClient.Get(h.ctx, canuseKey).Result()
	return IsNullKeyError(err)
}

func (h *RedisHandler) SetGlobalLock() {
	h.redisClient.Del(h.ctx, canuseKey).Result()
}

func (h *RedisHandler) SetGlobalUnLock() {
	h.redisClient.Set(h.ctx, canuseKey, canuseKey, 0).Result()
}
