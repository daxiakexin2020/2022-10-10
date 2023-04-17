package iface

type IConnection interface {
	Start()
	Stop()
	Connection() IConnection
	Server() IServer
	ConnId() uint64

	Send(data []byte) error
	SendToQueue(data []byte) error
	SendMsg(msgID uint32, data []byte) error     //直接将Message数据发送数据给远程的TCP客户端(无缓冲)
	SendBuffMsg(msgID uint32, data []byte) error //直接将Message数据发送给远程的TCP客户端(有缓冲)
}
