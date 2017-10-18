package goTCP

import (
	"errors"
	"fmt"
	"net"

	"strconv"
	"time"
)

type TCPService struct {
	Addr        *net.TCPAddr
	SessionList map[string]ISession
}
type IService interface {
	Start()
	CountSession() int
	GetList()
	Send(sessionId, val string) error
	SendAll(val string) error
	ShowMsg(sessionId, params string)
}

func NewService(port int) IService {
	addr, _ := net.ResolveTCPAddr("tcp", ":6000")
	self := &TCPService{
		Addr:        addr,
		SessionList: make(map[string]ISession),
	}

	return self

}

func (self *TCPService) SendAll(val string) error {
	var i int = 0
	if len(self.SessionList) > 0 {
		for item := range self.SessionList {
			i++
			self.SessionList[item].SendMsg(val)
		}
		return errors.New("发送终端数" + strconv.Itoa(i))
	}
	return errors.New("没有发现可用会话")
}

func (self *TCPService) Start() {

	listen, err := net.ListenTCP("tcp", self.Addr)

	if err != nil {
		fmt.Println("服务器关闭", err.Error())
	}

	for {
		//buf := make([]byte, 512)
		con, err := listen.AcceptTCP()

		if err != nil {
			fmt.Println("会话错误", err.Error())
			continue
		}
		//msg := string(buf[0:n])
		key := con.RemoteAddr().String()
		_, ok := self.SessionList[key]
		if ok {

			//session.ReceiveMsg(msg)
		} else {

			t := createSession(con.RemoteAddr(), con)
			//t.ReceiveMsg(msg)
			self.SessionList[key] = t
		}

	}

}

func (self *TCPService) CountSession() int {
	return len(self.SessionList)
}
func (self *TCPService) GetList() {

	if len(self.SessionList) == 0 {
		println("没有任何终端")
	}

	for session := range self.SessionList {
		self.SessionList[session].ToString()
	}
}
func (self *TCPService) ShowMsg(sessionId, params string) {
	session, ok := self.SessionList[sessionId]
	if ok {
		if params == "-r" {
			session.ShowReceiveMsg()
		}
		if params == "-s" {
			session.ShowSendMsg()
		}
	}
}
func (self *TCPService) Send(sessionId, val string) error {

	session, ok := self.SessionList[sessionId]
	if ok {
		return session.SendMsg(val)
	}
	return errors.New("没有发现可用会话")
}

type Session struct {
	Addr           net.Addr
	Key            string
	Service        *net.TCPConn
	ReceiveMsgList []string
	SendMsgList    []string
}

func createSession(addr net.Addr, conn *net.TCPConn) ISession {

	self := &Session{
		Service:        conn,
		Addr:           addr,
		Key:            addr.String(),
		ReceiveMsgList: make([]string, 0),
		SendMsgList:    make([]string, 0),
	}

	return self
}

type ISession interface {
	ToString()
	ReceiveMsg(val string)
	SendMsg(val string) error
	ShowReceiveMsg()
	ShowSendMsg()
}

func (self *Session) ShowReceiveMsg() {

	if len(self.ReceiveMsgList) == 0 {
		println("没有接受到任何消息")
	}

	for e := range self.ReceiveMsgList {
		fmt.Println(self.ReceiveMsgList[e])
	}

}
func (self *Session) ShowSendMsg() {
	if len(self.ReceiveMsgList) == 0 {
		println("没有发送任何消息")
	}
	for e := range self.SendMsgList {
		fmt.Println(self.SendMsgList[e])
	}
}

func format(tp, val string) string {
	t := time.Now()

	var str = "[ (" + tp + ")"
	year, month, day := t.Date()
	str += strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)
	str += " "
	hour, min, sec := t.Clock()
	str += strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
	str += "]" + val
	return str
}

func (self *Session) ReceiveMsg(val string) {

	self.ReceiveMsgList = append(self.ReceiveMsgList, format("接受", val))
}
func (self *Session) ToString() {

	fmt.Println("key:", self.Key)

}

func (self *Session) SendMsg(val string) error {
	self.SendMsgList = append(self.SendMsgList, format("发送", val))
	_, err := self.Service.Write([]byte(val))
	//_, err := self.Service.WriteToUDP([]byte(val), self.Addr)
	return err
}
