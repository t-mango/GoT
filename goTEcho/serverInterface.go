package goTEcho

import (
	"GoT/goTEvent"

	"golang.org/x/net/websocket"

	"github.com/labstack/echo"
)

//WebServiceInterface is interface
type WebServiceInterface interface {
	Start()
}

//WebService is service
type WebService struct {
	*echo.Echo
	ChanQueue goTEvent.EventQueue
}

//IWebSession
type IWebSession interface {
}
type WebSession struct {
	Guid string
	Conn *websocket.Conn
}
