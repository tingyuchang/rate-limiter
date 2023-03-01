package rate_limiter

import (
	"strconv"
	"time"
)

func TokenBucket(userId string, intervalSecond, maximumRequest int64) bool {
	now := time.Now()
	val, _ := rds.Get(ctx, userId+"_last_reset_time").Result()

	lastResetTime, _ := strconv.ParseInt(val, 10, 64)

	if now.Unix()-lastResetTime >= intervalSecond {
		// some problem here
		rds.Set(ctx, userId+"_count", strconv.FormatInt(maximumRequest, 10), 0)
	} else {
		val2, _ := rds.Get(ctx, userId+"_count").Result()
		requests, _ := strconv.ParseInt(val2, 10, 64)
		if requests <= 0 {
			return false
		}
	}

	rds.Decr(ctx, userId+"_count")

	return true
}
