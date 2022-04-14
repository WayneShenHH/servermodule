package redis

func gameRecordIncrScript() string {
	return `
	local key   = tostring(KEYS[1])
	local expireSec = tonumber(ARGV[1])
	
	if redis.call("EXISTS", key) == 0 then
		redis.call("SET", key, 1, "EX", expireSec)
		return 1
	end
	
	return redis.call("INCR", key)`
}

func updatePlayerProfitScript() string {
	return `
	local key = tostring(KEYS[1])
	local balance = tonumber(ARGV[1])

	redis.call("HINCRBY", key, "totalBalance", balance)
	redis.call("HSET", key, "lastBalance", balance)
	redis.call("EXPIRE", key, 14400)
	return 1`
}

func updateParadiseScript() string {
	return `
		local key = tostring(KEYS[1])
		local mj = tonumber(ARGV[1])
		local mi = tonumber(ARGV[2])
		local mo = tonumber(ARGV[3])
		local bj = tonumber(ARGV[4])
		local bi = tonumber(ARGV[5])
		local bo = tonumber(ARGV[6])

		redis.call("HINCRBY", key, "memberJuiceAmount", mj)
		redis.call("HINCRBY", key, "memberIncome", mi)
		redis.call("HINCRBY", key, "memberOutcome", mo)
		redis.call("HINCRBY", key, "botJuiceAmount", bj)
		redis.call("HINCRBY", key, "botIncome", bi)
		redis.call("HINCRBY", key, "botOutcome", bo)
		return 1`
}

func incrGameRoundScript() string {
	return `
		local key = tostring(KEYS[1])
		
		local gameRoundLocal = tonumber(redis.call("HGET",key,"gameRound")) + 1
		
		if (gameRoundLocal > tonumber(redis.call("HGET",key,"gameRoundLimit"))) then
			redis.call("HSET", key, "gameRound", 0)
			return 1
		end
		redis.call("HSET", key, "gameRound", gameRoundLocal)
		return 0
		`
}

func updateGame29ParadiseScript() string {
	return `
		local key = tostring(KEYS[1])
		local memberIncomeOfWheel = tonumber(ARGV[1])
		local memberOutcomeOfWheel = tonumber(ARGV[2])
		local memberJuiceAmountOfWheel = tonumber(ARGV[3])
		local memberIncomeOfBanker = tonumber(ARGV[4])
		local memberOutcomeOfBanker = tonumber(ARGV[5])
		local memberJuiceAmountOfBanker = tonumber(ARGV[6])

		redis.call("HINCRBY", key, "memberIncomeOfWheel", memberIncomeOfWheel)
		redis.call("HINCRBY", key, "memberOutcomeOfWheel", memberOutcomeOfWheel)
		redis.call("HINCRBY", key, "memberJuiceAmountOfWheel", memberJuiceAmountOfWheel)
		redis.call("HINCRBY", key, "memberIncomeOfBanker", memberIncomeOfBanker)
		redis.call("HINCRBY", key, "memberOutcomeOfBanker", memberOutcomeOfBanker)
		redis.call("HINCRBY", key, "memberJuiceAmountOfBanker", memberJuiceAmountOfBanker)

		return 1`
}

func updateTotalBetScript() string {
	return `
		local totalBetKey = tostring(KEYS[1])
		
		local betLionYellow = tonumber(ARGV[1]) 
		local betLionGreen = tonumber(ARGV[2]) 
		local betLionRed = tonumber(ARGV[3]) 
		local betPandaYellow = tonumber(ARGV[4]) 
		local betPandaGreen = tonumber(ARGV[5]) 
		local betPandaRed = tonumber(ARGV[6]) 
		local betMonkeyYellow = tonumber(ARGV[7]) 
		local betMonkeyGreen = tonumber(ARGV[8]) 
		local betMonkeyRed = tonumber(ARGV[9]) 
		local betRabbitYellow = tonumber(ARGV[10]) 
		local betRabbitGreen = tonumber(ARGV[11]) 
		local betRabbitRed = tonumber(ARGV[12]) 
		local betBanker = tonumber(ARGV[13]) 
		local betPlayer = tonumber(ARGV[14]) 
		local betPush = tonumber(ARGV[15]) 

		if betLionYellow  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetLionYellow", betLionYellow)
		end
		if betLionGreen  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetLionGreen", betLionGreen)
		end
		if betLionRed  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetLionRed", betLionRed)
		end
		if betPandaYellow  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetPandaYellow", betPandaYellow)
		end
		if betPandaGreen  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetPandaGreen", betPandaGreen)
		end
		if betPandaRed  > 0 then
		redis.call("HINCRBY", totalBetKey, "totalBetPandaRed", betPandaRed)
		end
		if betMonkeyYellow  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetMonkeyYellow", betMonkeyYellow)
		end
		if betMonkeyGreen  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetMonkeyGreen", betMonkeyGreen)
		end
		if betMonkeyRed  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetMonkeyRed", betMonkeyRed)
		end
		if betRabbitYellow  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetRabbitYellow", betRabbitYellow)
		end
		if betRabbitGreen  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetRabbitGreen", betRabbitGreen)
		end
		if betRabbitRed  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetRabbitRed", betRabbitRed)
		end
		if betBanker  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetBanker", betBanker)
		end
		if betPlayer  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetPlayer", betPlayer)
		end
		if betPush  > 0 then
			redis.call("HINCRBY", totalBetKey, "totalBetPush", betPush)
		end

		return 1`
}

// TODO: 需要測試確保效能
func addSettlementDataScript() string {
	return `
		local skey = tostring(ARGV[1])
		local unskey = tostring(ARGV[2])
		local data = tostring(ARGV[3])
		local timestamp = tonumber(ARGV[4])
		local gameRecordId = tostring(ARGV[5])
		local startIndex = tonumber(ARGV[6])
		local stopIndex = tonumber(ARGV[7])
		
		redis.call("HSET", skey, gameRecordId, data)
		redis.call("ZADD", unskey, timestamp, gameRecordId)

		local members = redis.call("ZREVRANGE", unskey, startIndex, stopIndex)

		for _, member in ipairs(members) do
			redis.call("HDEL", skey, member)
			redis.call("ZREM", unskey, member)
		end
	
		return 1
	`
}
