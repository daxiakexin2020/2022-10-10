package service

import (
	"01/service/backend"
	"fmt"
	"reflect"
	"strings"
)

type Service interface {
	InitService() error
	Name() string
}

type Server struct {
	ReServices []Service
	AH         *backend.A     `rtag:"ah"`
	BH         *backend.B     `rtag:"bh"`
	CH         *backend.C     `rtag:"ch"`
	ConH       backend.Config `rtag:"conh"`
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() error {
	return s.run()
}

func (s *Server) RegisterService(service ...Service) {
	s.ReServices = append(s.ReServices, service...)
}

func (s *Server) run() error {

	rsType := reflect.TypeOf(s)
	rsValue := reflect.ValueOf(s)

	for _, rs := range s.ReServices {
		if err := rs.InitService(); err != nil {
			return err
		}
		var isStartOk bool
		for i := 0; i < rsType.Elem().NumField(); i++ {
			vFiled := rsType.Elem().Field(i)
			rsValueIndex := rsValue.Elem().Field(i)
			if strings.ToLower(rs.Name()) == strings.ToLower(vFiled.Tag.Get("rtag")) {
				if rsValueIndex.Kind() == reflect.Ptr {
					rsValueIndex.Set(reflect.ValueOf(rs)) // 指针，直接挂载服务
				} else {
					rsValueIndex.Set(reflect.ValueOf(rs).Elem()) // 非指针，解开引用,挂载服务
				}
				isStartOk = true
				break
			}
		}
		if !isStartOk {
			return fmt.Errorf("*****************%s服务启动失败，未找到相应的rtag:【%s】，请检查*********************\n", rs.Name(), rs.Name())
		}
	}
	fmt.Printf("*****************所有服务启动成功，一共启动【%d】个服务*********************\n", len(s.ReServices))
	return nil
}
