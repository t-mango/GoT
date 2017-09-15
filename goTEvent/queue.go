package goTEvent

type EventQueue interface {
	PostData(interface{})
}

type evQueue struct {
	Queue chan interface{}
}

func (self *evQueue) PostData(data interface{}) {
	self.Queue <- data
}

func NewEventQueue(num int) *evQueue {
	self := &evQueue{
		Queue: make(chan interface{}, num),
	}
	return self
}
