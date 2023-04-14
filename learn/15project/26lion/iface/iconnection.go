package iface

type IConnection interface {
	Start()
	Stop()
	Connection() IConnection
	Server() IServer
	ConnId() uint64
}
