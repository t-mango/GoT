package wtiot

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"golang.org/x/net/websocket"
)

const (
	WT_DEVICELIST_S     = "WT:DEVICELIST"
	WT_DEVICE_H         = "WT:DEVICE:"        //
	WT_DEVICECOMMAND_K  = "WT:DEVICECOMMAND:" //
	WT_DEVICE_HISTORY_H = "WT:DEVICEHISTORY:"
	WT_DEVICE_TIME_L    = "WT:DEVICETIME:"
	SubKey              = "WTCLIENTCHAN"
)

//注册redis
func hello(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				c.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username == "jon" && password == "shhh!" {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

func restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}
func Start() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	g := e.Group("/websocket")

	g.Use(middleware.BasicAuth(func(username, password string, coontext echo.Context) (bool, error) {

		if username == "mango" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))
	g.GET("/ws", hello)

	// e.POST("/login", login)

	initRoute(e)
	e.Static("/", "wtiot/html/index.html")

	// //jwt
	// e.GET("/", accessible)
	// r := e.Group("/restricted")
	// r.Use(middleware.JWT([]byte("secret")))
	// r.GET("", restricted)

	e.Logger.Fatal(e.Start(":80"))
}
