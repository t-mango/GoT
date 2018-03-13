package wtiot

import (
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "192.168.1.106:6379",
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

func GetDevices() ([]string, error) {
	return redisClient.SMembers(WT_DEVICELIST_S).Result()
}

func GetHistroy(deviceId string) ([]string, error) {

	sort := &redis.Sort{}
	sort.By = WT_DEVICE_HISTORY_H + deviceId + ":*"
	sort.Order = "DESC"
	sort.Get = []string{WT_DEVICE_HISTORY_H + deviceId + ":*->action", WT_DEVICE_HISTORY_H + deviceId + ":*->type", WT_DEVICE_HISTORY_H + deviceId + ":*->state"} //, "type", "content", "time"} //, "type", "content"} //, "state", "state_0", "state_1"}
	return redisClient.Sort(WT_DEVICE_TIME_L+deviceId, sort).Result()

	// if err != nil {

	// } else {

	// }

	// for _, value := range list {
	// 	historys["1"] = value.(string)
	// }
	// return nil, historys
}
