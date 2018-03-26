package wtiot

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type ResParam struct {
	Result int
	Msg    string
	Data   interface{}
}

func deviceList(c echo.Context) error {
	req := c.Request()
	fmt.Println("Proto", req.Proto)

	list, err := GetDevices()
	if err != nil {
		list = []string{}
	}
	result := new(ResParam)
	result.Result = 0
	result.Msg = "success"
	result.Data = list
	return c.JSON(http.StatusOK, result)
}

func deviceCmdhistorylist(c echo.Context) error {
	deviceId := c.Param("divicerId")
	result := new(ResParam)
	list, err := GetHistroy(deviceId)
	topArray := make([][]string, 0)

	length := len(list)
	temp := 10
	for i := 0; temp*i+(temp-1) <= length; i++ {

		topArray = append(topArray, list[temp*i:temp*i+temp])
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
		temp["keyTime"] = item[8]
		temp["content"] = item[9]
		topMap = append(topMap, temp)
	}
	result.Data = topMap

	if err != nil {
		if err != nil {
			result.Result = 1
			result.Msg = "命令下达失败:" + err.Error()
		}
	}

	return c.JSON(http.StatusOK, result)
}

func cmdAction(c echo.Context) error {
	deviceId := c.Param("divicerId")
	cmd := c.Param("action")
	result := new(ResParam)
	if deviceId != "" && cmd != "" {

		_, err := ActionCommand(deviceId, cmd)
		if err != nil {
			result.Result = 1
			result.Msg = "命令下达失败:" + err.Error()
		} else {
			result.Result = 0
			result.Msg = "success"
		}
	} else {
		result.Result = 1
		result.Msg = "命令下达失败:参数错误"
	}
	return c.JSON(http.StatusOK, result)
	// return c.String(http.StatusOK, "Hello, World!")
}
