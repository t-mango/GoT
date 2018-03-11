package wtiot

import (
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.106:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func ZAdd(x redis.Z) (int64, error) {

	return redisClient.ZAdd("devices:list", x).Result()
}
