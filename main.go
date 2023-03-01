package rate_limiter

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var rds *redis.Client
var ctx = context.Background()

func main() {
	rds = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
