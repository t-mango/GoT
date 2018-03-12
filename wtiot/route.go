package wtiot

import "github.com/labstack/echo"

func initRoute(e *echo.Echo) {

	///device/list
	e.GET("/device/list", deviceList)

	///device/cmdhistorylist
	e.GET("/device/cmdhistorylist", deviceCmdhistorylist)
	//device/list

	e.GET("/:divicerId/:action", cmdAction)
}
