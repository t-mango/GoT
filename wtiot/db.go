package wtiot

import (
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

const SubKey = "wtClientChan"

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.5.71:6379",
		Password: "", // no password set
		DB:       2,  // use default DB
	})
}

func ZAdd(x redis.Z) (int64, error) {

	return redisClient.ZAdd("devices:list", x).Result()
}

func ActionCommand(deviceId, cmd string) (int64, error) {

	return redisClient.Publish(SubKey, deviceId+"|"+cmd).Result()
}
