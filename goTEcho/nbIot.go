package goTEcho

import (
	"github.com/labstack/echo"
	"golang.org/x/net/websocket"
)

func (self *WebService) wsNBiot(c echo.Context) error {

	guid := c.Param("id")

	websocket.Handler(func(ws *websocket.Conn) {

		session := &WebSession{
			Guid: guid,
			Conn: ws,
		}
		self.ChanQueue.PostData(session)
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
