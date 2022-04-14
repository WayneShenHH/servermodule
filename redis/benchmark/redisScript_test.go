package benchmark

func getScript() string {
	return `
	local key   = tostring(KEYS[1])
	local expireSec = tonumber(ARGV[1])
	
	if redis.call("EXISTS", key) == 0 then
		redis.call("SET", key, 1, "EX", expireSec)
		return 1
	end
	
	return redis.call("INCR", key)`
}

func setScript() string {
	return `
	local key = tostring(KEYS[1])
	local expireSec = tonumber(ARGV[1])
	local ii = tonumber(ARGV[2])
	local iii = tonumber(ARGV[3])
	
	if redis.call("EXISTS", key) == 0 then
		redis.call("HSET", key, "totalBalance", ii)
		redis.call("HSET", key, "lastBalance", iii)
		return 1
	end

	redis.call("HSET", key, "totalBalance", ii)
	redis.call("HSET", key, "lastBalance", iii)
	return 1`
}
