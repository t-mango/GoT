package wtiot

import (
	"GoT/goTEvent"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

var wsTask goTEvent.ITask = goTEvent.CreateTask()

type ConnInfo struct {
	DeviceId string
	PageId   string

	Conn *websocket.Conn
}

var deviceMap = make(map[string][]*websocket.Conn)

var WT_DEVICE_CHANGE = "WT:DEV:CHANGE"

//注册redis
func hello(c echo.Context) error {
	deviceId := c.Param("divicerId")
	pageId := c.Param("pageId")

	websocket.Handler(func(ws *websocket.Conn) {
		wsTask.PushData(ConnInfo{DeviceId: deviceId, PageId: pageId, Conn: ws})
		defer ws.Close()
		deviceSessionList, ok := deviceMap[deviceId]
		if !ok {
			deviceSessionList = make([]*websocket.Conn, 0)

		}
		deviceSessionList = append(deviceSessionList, ws)
		deviceMap[deviceId] = deviceSessionList
		data := Respdata{}
		deviceInfo, err := HGetAll(WT_DEVICE_H + deviceId)
		if err != nil {
			data.DeviceInfo = err.Error()
		} else {
			data.DeviceInfo = deviceInfo
		}
		json, _ := json.Marshal(data)
		for {

			//Write
			err := websocket.Message.Send(ws, string(json))
			if err != nil {
				c.Logger().Error(err)
				break

			}
			// // Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				fmt.Println("链接错误")
				c.Logger().Error(err)
				break
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

type Respdata struct {
	DeviceInfo interface{}
	Cmdhistory interface{}
}

func subThead() {
	subHandler := redisClient.Subscribe(WT_DEVICE_CHANGE)
	for {
		subData, err := subHandler.ReceiveMessage()
		if err != nil {
			fmt.Println("错误", err.Error())
			continue
		}
		fmt.Println("subThead", subData.Payload)

		deviceId := strings.Split(subData.Payload, "|")[0]
		timestamp := strings.Split(subData.Payload, "|")[1]
		list, ok := deviceMap[deviceId]
		//整理数据

		data := Respdata{}

		history, err := HGetAll(WT_DEVICE_HISTORY_H + deviceId + ":" + timestamp)
		if err != nil {
			data.Cmdhistory = err.Error()
		} else {
			data.Cmdhistory = history
		}
		deviceInfo, err := HGetAll(WT_DEVICE_H + deviceId)
		if err != nil {
			data.DeviceInfo = err.Error()
		} else {
			data.DeviceInfo = deviceInfo
		}
		json, _ := json.Marshal(data)
		if ok {
			for _, item := range list {
				err := websocket.Message.Send(item, string(json))
				if err != nil {
					delete(deviceMap, deviceId)
				}
			}
		}

	}
}

// func messageData(deviceId, timestamp string) string {

// 	jsonMap := make(map[string]string)

// }

func Start() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	g := e.Group("/websocket")
	go subThead()
	// g.Use(middleware.BasicAuth(func(username, password string, coontext echo.Context) (bool, error) {

	// 	if username == "mango" && password == "secret" {
	// 		return true, nil
	// 	}
	// 	return false, nil
	// }))
	g.GET("/:divicerId", hello)

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
