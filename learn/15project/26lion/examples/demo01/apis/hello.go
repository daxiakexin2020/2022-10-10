package apis

import (
	"26lion/iface"
	"26lion/net"
	"log"
)

type Hello struct {
	net.BaseRouter
}

func (h *Hello) Handle(req iface.IRequest) {
	err := req.Connection().Send([]byte("Hello Api>>>>>>>"))
	if err != nil {
		log.Println("Hello Api Send Error :", err)
	} else {
		log.Println("Hello Api Send Ok.........")
	}
}

func NewHell() *Hello {
	return &Hello{}
}
