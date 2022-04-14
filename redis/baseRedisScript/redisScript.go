package baseRedisScript

import "github.com/go-redis/redis/v8"

func CheckDel() *redis.Script {
	return redis.NewScript(`
		if redis.call("EXISTS", KEYS[1]) == 1 then
			local res = redis.call("GET", KEYS[1])
			if res == ARGV[1] then
				return redis.call("DEL", KEYS[1])
			end
			return "mismatch"
		end
		return false
	`)
}

func TryLock() *redis.Script {
	return redis.NewScript(`
		if redis.call("EXISTS", KEYS[1]) == 0 then
			redis.call("SET", KEYS[1], ARGV[1], "EX", 120)
			return 1
		end
		return 0
	`)
}
