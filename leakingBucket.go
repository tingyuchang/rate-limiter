package rate_limiter

func LeakingBucket(userId, requestId string, intervalSecond, maximumRequest int64) bool {
	requestCount := rds.LLen(ctx, userId).Val()
	if requestCount > maximumRequest {
		return false
	}

	rds.RPush(ctx, userId, requestId)

	return true
}
