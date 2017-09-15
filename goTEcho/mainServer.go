package goTEcho

import "GoT/goTEvent"
import "github.com/labstack/echo"

func NewWebService(e *echo.Echo, queue goTEvent.EventQueue) *WebService {

	self := &WebService{
		Echo:      e,
		ChanQueue: queue,
	}

	return self
}

func (self *WebService) Start() {

	self.initRoute() //初始化路由

}
