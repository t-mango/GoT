package wtiot

import (
	"net/http"

	"github.com/labstack/echo"
)

type ResParam struct {
	Result int
	Msg    string
	Data   interface{}
}

func deviceList(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}

func deviceCmdhistorylist(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
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
