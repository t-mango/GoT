package goTEvent

type ITask interface {
	GetTask() chan interface{}
	PushData(interface{})
}

type chanTask struct {
	Allot chan interface{}
}

func (self *chanTask) GetTask() chan interface{} {
	return self.Allot
}

func (self *chanTask) PushData(anyData interface{}) {
	self.Allot <- anyData
}

func CreateTask() ITask {

	self := &chanTask{
		Allot: make(chan interface{}, 1000),
	}
	return self
}
