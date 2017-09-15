package main

import (
	"GoT/goTEcho"
	"GoT/goTEvent"
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func test1(c echo.Context) error {

	fmt.Println(1, c.FormValue("id"))
	//fmt.Println(2, c.FormParams("id"))
	fmt.Println(3, c.Param("id"))
	fmt.Println(4, c.QueryParam("id"))

	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		i := 0
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!,"+strconv.Itoa(i))
			if err != nil {

				c.Logger().Error(err)
				return
			}
			fmt.Println("到了", i)
			// // Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
				return
			}
			fmt.Printf("%s\n", msg)
			i++
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

func main() {

	//创建echo 数据
	e := echo.New()
	query := goTEvent.NewEventQueue(50)

	server := goTEcho.NewWebService(e, query)
	server.Start()

	  <- query.Queue

	// e.Static("/", "public/index.html")
	// e.GET("/dd/:id", test1)
	e.Logger.Fatal(e.Start(":8000"))

	//fmt.Println("")
	// ssh.Start()
	// service := ":8888"
	// IPAddr, err := net.ResolveIPAddr("ip4", service)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// IPconn, err := net.ListenIP("ip:tcp", IPAddr)

	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }

	// for {
	// 	// listen2Client function will listen forever until terminate
	// 	// up to you to define how listen2Client works...
	// 	byt := make([]byte, 1024)
	// 	index, _ := IPconn.Read(byt)

	// 	fmt.Println(string(byt[0:index]))
	// }
}
