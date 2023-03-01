package rate_limiter

import (
	"strconv"
	"time"
)

func RateLimitWithFixedWindow(usedID string, intervalSecond int64, maximumRequests int64) bool {
	now := time.Now()
	currentWindow := strconv.FormatInt(now.Unix()/intervalSecond, 10)

	key := usedID + ":" + currentWindow

	value, _ := rds.Get(ctx, key).Result()
	requestCount, _ := strconv.ParseInt(value, 10, 64)

	if requestCount > maximumRequests {
		// drop request
		return false
	}

	rds.Incr(ctx, key)
	rds.Expire(ctx, key, 120*time.Second)

	return true

}
