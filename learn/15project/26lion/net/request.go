package net

import (
	"26lion/iface"
	"sync"
)

const (
	PRE_HANDLE iface.HandleStep = iota
	HANDLE
	POST_HANDLE

	HANDLE_OVER
)

type Request struct {
	conn     iface.IConnection
	msg      iface.IMessage
	router   iface.IRouter
	steps    iface.HandleStep
	stepLock *sync.RWMutex
	needNext bool
}

var _ iface.IRequest = (*Request)(nil)

func NewRequest(msg iface.IMessage, conn iface.IConnection) *Request {
	return &Request{
		conn:     conn,
		msg:      msg,
		router:   nil,
		steps:    PRE_HANDLE,
		needNext: true,
	}
}

func (r *Request) Message() iface.IMessage {
	return r.msg
}

func (r *Request) Connection() iface.IConnection {
	return r.conn
}

func (r *Request) Data() []byte {
	return r.msg.GetData()
}

func (r *Request) MsgId() uint32 {
	return r.msg.GetMsgID()
}

func (r *Request) BindRouter(router iface.IRouter) {
	r.router = router
}

func (r *Request) Call() {
	if r.router == nil {
		return
	}
	for r.steps < HANDLE_OVER {
		switch r.steps {
		case PRE_HANDLE:
			r.router.PreHandle(r)
		case HANDLE:
			r.router.Handle(r)
		case POST_HANDLE:
			r.router.PostHandle(r)
		}
		r.next()
	}
	r.steps = PRE_HANDLE
}

func (r *Request) next() {
	if !r.needNext {
		r.needNext = true
		return
	}
	r.stepLock.Lock()
	r.steps++
	r.stepLock.Unlock()
}

func (r *Request) Abort() {
	r.stepLock.Lock()
	r.steps = HANDLE_OVER
	r.stepLock.Unlock()
}
