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

func TestActionCommand(t *testing.T) {

	_, err := wtiot.ActionCommand("36cfeaa4c23047ad8ab4c5f9a7a79ec6", "2")
	if err != nil {
		t.Error("错误", err.Error())
	}

}
func TestGetDevices(t *testing.T) {

	strlist, err := wtiot.GetDevices()
	if err != nil {
		t.Error("错误", err.Error())
	}

	for item := range strlist {
		t.Log(item)
		//fmt.Println(item)
	}

}
