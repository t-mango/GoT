package wtiot

import "github.com/labstack/echo"

func initRoute(e *echo.Echo) {

	///device/list
	e.GET("/device/list", deviceList)

	///device/cmdhistorylist
	e.GET("/device/cmdhistorylist/:divicerId", deviceCmdhistorylist)
	//device/list

	e.GET("/device/action/:divicerId/:action", cmdAction)
}
