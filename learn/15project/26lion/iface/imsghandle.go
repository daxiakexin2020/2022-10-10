package iface

type IMsgHandle interface {
	AddRouter(msgID uint32, r IRouter)
	StartWorkerPool()
	StopWorkerPool()
	SendMsgToTaskQueue(request IRequest)
	Execute(request IRequest)
}
