package rate_limiter

import (
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

func SlidingLogs(userId, requestId string, intervalSecond int64, maximumRequests int64) bool {
	now := time.Now()
	currentTime := strconv.FormatInt(now.Unix(), 10)
	lastWindowTime := strconv.FormatInt(now.Unix()-intervalSecond, 10)
	// zcount is O(log(N))
	requestCount := rds.ZCount(ctx, userId, lastWindowTime, currentTime).Val()
	if requestCount > maximumRequests {
		return false
	}

	rds.ZAdd(ctx, userId, redis.Z{
		Score:  float64(now.Unix()),
		Member: requestId,
	})
	return true
}
