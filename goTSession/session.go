package goTSession

// 管理 websocket session
// 管理 upd session
// 管理 tcp session

type sessionManger struct {
	ListWeb map[string]ISession
}
type ISession interface {
	Send(interface{})
}
