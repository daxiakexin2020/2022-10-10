package iface

type HandleStep int

type IRequest interface {
	Connection() IConnection
	Data() []byte
	MsgId() uint32
	BindRouter(r IRouter)
	Call()
}
