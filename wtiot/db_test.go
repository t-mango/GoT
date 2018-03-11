package wtiot_test

import (
	"GoT/wtiot"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestAdd(t *testing.T) {

	x := redis.Z{}
	x.Member = "99594383ff9e1b03dd1ba6e2170b3282"
	x.Score, _ = strconv.ParseFloat(strconv.FormatInt(time.Now().Unix(), 10), 64)
	_, err := wtiot.ZAdd(x)

	if err != nil {
		t.Error("错误", err.Error())
	}

}
