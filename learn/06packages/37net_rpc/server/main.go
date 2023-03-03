package main

import (
	"37net_rpc/proto"
	"errors"
	"net"
	"net/http"
	"net/rpc"
)

type Server struct {
}

var ps *proto.PServer

func (s *Server) Fetch(id string, replay *string) error {
	if p, ok := ps.List[id]; ok {
		*replay = p.Name
		return nil
	}
	return errors.New("no this people")
}

func main() {
	initMockData()
	rpc.Register(&Server{})
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":9112")
	if e != nil {
		panic(e)
	}
	http.Serve(l, nil)
}

func initMockData() {
	ps = &proto.PServer{List: map[string]*proto.People{}}
	ps.List["1"] = &proto.People{Name: "zz", Age: 320}
}
