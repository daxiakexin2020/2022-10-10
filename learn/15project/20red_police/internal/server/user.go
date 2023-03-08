package server

import (
	"errors"
	"fmt"
)

type RegisterRequest struct {
	Name string `json:"name" mapstructure:"name"`
	//Pwd   string `json:"pwd"  mapstructure:"pwd"`
	//Phone string `json:"phone"  mapstructure:"phone"`
}

func (s *Server) Register(req *RegisterRequest, res *RegisterRequest) error {
	fmt.Println("我被调用了", req)
	res = &RegisterRequest{
		Name: "ss",
	}
	return errors.New("test err")
}

func (s *Server) Register2(req *string, res *int) error {
	fmt.Println("我被调用了2", req)
	b := 3
	*res = b
	return errors.New("test err")
}
