package rate_limiter

import (
	"strconv"
	"time"
)

func SlidingWindow(userId, requestId string, intervalSecond, maximumRequest int64) bool {
	now := time.Now()
	currentWindow := strconv.FormatInt(now.Unix()/intervalSecond, 10)

	key := userId + ":" + currentWindow

	val, _ := rds.Get(ctx, key).Result()
	currentRequests, _ := strconv.ParseInt(val, 10, 64)
	if currentRequests >= maximumRequest {
		return false
	}

	lastWindow := strconv.FormatInt((now.Unix()-intervalSecond)/intervalSecond, 10)
	lastKey := userId + ":" + lastWindow
	val, _ = rds.Get(ctx, lastKey).Result()
	lastRequests, _ := strconv.ParseInt(val, 10, 64)

	elapsedTimePercentage := float64(now.Unix()%intervalSecond) / float64(intervalSecond)

	if (float64(lastRequests)*(1-elapsedTimePercentage))+float64(lastRequests) >= float64(maximumRequest) {
		// drop request
		return false
	}

	rds.Incr(ctx, key)

	return true
}
