package server

import (
	"13table_driver/service"
	"fmt"
)

type Server struct {
	Srv service.Servicer
}

// todo 问题：一个类，其实就是单例
var services map[int]service.Servicer = map[int]service.Servicer{
	1: service.NewA(),
	2: service.NewB(),
	3: service.NewC(),
}

// 使用自定义函数，可以解决单例问题
var fservices map[int]f = map[int]f{
	1: func() service.Servicer {
		return service.NewA()
	},
	2: func() service.Servicer {
		return service.NewB()
	},
	3: func() service.Servicer {
		return service.NewC()
	},
}

type f func() service.Servicer

func NewServe(flag int) (*Server, error) {
	if service, ok := fservices[flag]; ok {
		return &Server{Srv: service()}, nil
	}
	return nil, fmt.Errorf("暂不支持此中策略:=d", flag)

	if service, ok := services[flag]; ok {
		return &Server{Srv: service}, nil
	}
	return nil, fmt.Errorf("暂不支持此中策略:=d", flag)
}
