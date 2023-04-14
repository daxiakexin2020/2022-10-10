package net

import (
	"26lion/config"
	"26lion/iface"
	"context"
	"fmt"
	"log"
)

type MsgHandle struct {
	apis           map[uint32]iface.IRouter
	workerPoolSize uint32
	taskQueue      []chan iface.IRequest
	ctx            context.Context
	cancel         context.CancelFunc
	isRunning      bool
}

func NewMagHandle() *MsgHandle {
	return &MsgHandle{
		apis:           map[uint32]iface.IRouter{},
		workerPoolSize: config.GetMsgHandleConfig().WorkPoolSize,
		taskQueue:      make([]chan iface.IRequest, config.GetMsgHandleConfig().WorkPoolSize),
	}
}

func (msg *MsgHandle) Execute(request iface.IRequest) {
	if msg.isRunning {
		msg.SendMsgToTaskQueue(request)
	} else {
		go msg.do(request)
	}
}

func (msg *MsgHandle) AddRouter(msgID uint32, r iface.IRouter) {
	if _, ok := msg.apis[msgID]; ok {
		msgErr := fmt.Sprintf("repeated api , msgID = %+v\n", msgID)
		panic(msgErr)
	}
	msg.apis[msgID] = r
}

func (msg *MsgHandle) SendMsgToTaskQueue(request iface.IRequest) {
	if msg.workerPoolSize <= 0 || !msg.isRunning {
		return
	}
	workerID := request.Connection().ConnId() % uint64(msg.workerPoolSize)
	msg.taskQueue[workerID] <- request
}

func (msg *MsgHandle) StartWorkerPool() {
	if msg.workerPoolSize > 0 {
		ctx, cancelFunc := context.WithCancel(context.Background())
		msg.ctx = ctx
		msg.cancel = cancelFunc
		for i := 0; i < int(msg.workerPoolSize); i++ {
			msg.taskQueue[i] = make(chan iface.IRequest, config.GetMsgHandleConfig().WorkPoolSize)
			go msg.startWorker(i, msg.taskQueue[i])
		}
		msg.isRunning = true
	}
}

func (msg *MsgHandle) startWorker(workerID int, taskQueue <-chan iface.IRequest) {
	log.Printf("Worker ID = %d is started.\n", workerID)
	for {
		select {
		case req := <-taskQueue:
			msg.do(req)
		case <-msg.ctx.Done():
			log.Printf("worker ID=%d is stop.......\n", workerID)
			return
		default:

		}
	}
}

func (msg *MsgHandle) do(request iface.IRequest) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("do MsgHandler panic: %v", err)
		}
	}()
	router, ok := msg.apis[request.MsgId()]
	if !ok {
		log.Printf("api msgID = %d is not FOUND!", request.MsgId())
		return
	}
	request.BindRouter(router)
	request.Call()
}

func (msg *MsgHandle) StopWorkerPool() {
	if !msg.isRunning {
		return
	}
	msg.cancel()
	msg.isRunning = false
}
