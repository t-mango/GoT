package wtiot

import (
	"net/http"

	"github.com/labstack/echo"
)

func deviceList(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}

func deviceCmdhistorylist(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func cmdAction(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
