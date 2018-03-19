package wtiot_test

import (
	"GoT/wtiot"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis"
)

func TestAdd(t *testing.T) {

	x := redis.Z{}
	x.Member = "66cfeaa4c23047ad8ab4c5f9a7a79ec6"
	x.Score, _ = strconv.ParseFloat(strconv.FormatInt(time.Now().Unix(), 10), 64)
	_, err := wtiot.ZAdd(x)

	if err != nil {
		t.Error("错误", err.Error())
	}

}
func TestGetHistroy(t *testing.T) {

	list, err := wtiot.GetHistroy("36cfeaa4c23047ad8ab4c5f9a7a79ec6")
	if err != nil {
		t.Error("错误", err.Error())
	}
	if !(len(list) > 0) {
		t.Error("没有数据")
	}

	topArray := make([][]string, 0)

	length := len(list)
	for i := 0; 8*i+7 <= length; i++ {

		topArray = append(topArray, list[8*i:8*i+8])
	}

	topMap := make([]map[string]string, 0)
	for _, item := range topArray {
		temp := make(map[string]string)
		temp["time"] = item[0]
		temp["action"] = item[1]
		temp["type"] = item[2]
		temp["state"] = item[3]
		temp["state_0"] = item[4]
		temp["state_1"] = item[5]
		temp["state_2"] = item[6]
		temp["state_3"] = item[7]
		topMap = append(topMap, temp)
	}

	fmt.Println(topMap)

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
func TestHGetAll(t *testing.T) {

	strlist, err := wtiot.HGetAll(wtiot.WT_DEVICE_HISTORY_H+"36cfeaa4c23047ad8ab4c5f9a7a79ec6:1521013941")

	if err != nil {
		t.Error("错误", err.Error())
	}

	for item := range strlist {
		t.Log(item)
		//fmt.Println(item)
	}

}
