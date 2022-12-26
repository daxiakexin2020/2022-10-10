package server_govern

import (
	"14gateway/register_center"
	"errors"
	"fmt"
	"strings"
)

type server_type string

type method_type string

type register_type string

const (
	HTTP server_type = "http"
	GRPC server_type = "grpc"
)

const (
	ETCD   register_type = "etcd"
	CONSOL register_type = "consol"
)

const (
	GET     method_type = "GET"
	POST    method_type = "POST"
	PUT     method_type = "PUT"
	DELETE  method_type = "DELETE"
	OPTIONS method_type = "OPTIONS"
	HEAD    method_type = "HEAD"
	ANY     method_type = "ANY"
)

type server struct {
	name  string        `json:"name"`
	addr  []string      `json:"addr"`
	stype server_type   `json:"stype"`
	mtype method_type   `json:"mtype"`
	rtype register_type `json:"rtype"`
}

type Option func(s *server)

func NewServer(name string, option ...Option) *server {
	s := &server{
		name:  name,
		stype: HTTP,
		mtype: GET,
		rtype: ETCD,
	}
	s.apply(option)
	return s
}

func (s *server) apply(opts []Option) {
	for _, opt := range opts {
		opt(s)
	}
}

func WithAddr(addr []string) Option {
	return func(s *server) {
		s.addr = addr
	}
}

func WithServerType(serverType server_type) Option {
	return func(s *server) {
		s.stype = serverType
	}
}

func WithMethodType(methodType method_type) Option {
	return func(s *server) {
		s.mtype = methodType
	}
}

func WithRegisterType(registerType register_type) Option {
	return func(s *server) {
		s.rtype = registerType
	}
}

func (s *server) getRegisteServer() (*register_center.RegisteServer, error) {
	switch s.rtype {
	case ETCD:
		return register_center.NewRegisteServer(register_center.ETCD)
	case CONSOL:
		return register_center.NewRegisteServer(register_center.CONSOL)
	default:
		return nil, errors.New("没有注册中心可以使用")
	}
}

func (s *server) makeKey() string {
	return fmt.Sprintf("%s.%s.%s", s.name, s.stype, s.mtype)
}

func (s *server) makeValue() string {
	return strings.Join(s.addr, ",")
}

func GetMethod(method string) method_type {
	switch method_type(method) {
	case GET:
		return GET
	case POST:
		return POST
	case PUT:
		return PUT
	case HEAD:
		return HEAD
	case DELETE:
		return DELETE
	case OPTIONS:
		return OPTIONS
	case ANY:
		return ANY
	default:
		return ""
	}
}
