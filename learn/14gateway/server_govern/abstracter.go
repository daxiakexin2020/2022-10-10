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

func GetMethod(method string) method_type {
	switch method_type(method) {
	case GET_TYPE:
		return GET_TYPE
	case POST_TYPE:
		return POST_TYPE
	case PUT_TYPE:
		return PUT_TYPE
	case HEAD_TYPE:
		return HEAD_TYPE
	case DELETE_TYPE:
		return DELETE_TYPE
	case OPTIONS_TYPE:
		return OPTIONS_TYPE
	default:
		return ""
	}
}
