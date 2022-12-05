package server_govern

import (
	"14gateway/register_center"
	"fmt"
	"strings"
)

type server_type string

type method_type string

const (
	HTTP_TYPE    server_type = "http"
	GRPC_TYPE    server_type = "grpc"
	GET_TYPE                 = "GET"
	POST_TYPE                = "POST"
	PUT_TYPE                 = "PUT"
	DELETE_TYPE              = "DELETE"
	OPTIONS_TYPE             = "OPTIONS"
	HEAD_TYPE                = "HEAD"
	ANY_TYPE                 = "ANY"
)

type server struct {
	Name  string      `json:"name"`
	Addr  []string    `json:"addr"`
	Stype server_type `json:"stype"`
	Mtype method_type `json:"mtype"`
}

func NewServer(name string, addr []string, stype server_type, mtype method_type) *server {
	return &server{
		Name:  name,
		Addr:  addr,
		Stype: stype,
		Mtype: mtype,
	}
}

func (s *server) getRegisteServer() (*register_center.RegisteServer, error) {
	return register_center.NewRegisteServer(register_center.ETCD)
}

func (s *server) makeKey() string {
	return fmt.Sprintf("%s.%s.%s", s.Name, s.Stype, s.Mtype)
}

func (s *server) makeValue() string {
	return strings.Join(s.Addr, ",")
}
