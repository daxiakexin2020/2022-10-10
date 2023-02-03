package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"service/service"
)

type Demo struct {
}

func (d *Demo) Test(ctx context.Context, in *service.UserReuqest) (*service.UserResponse, error) {
	fmt.Println(ctx.Value("test_ctx_key"), in.Id)
	return &service.UserResponse{
		Name: "TEST1",
	}, nil
}

func (d *Demo) Test2(ctx context.Context, in *service.UserReuqest) (*service.UserResponse, error) {
	return &service.UserResponse{
		Name: "Test2",
	}, nil
}

func main() {
	HandleTest()
}

func HandleTest() {

	srv := grpc.NewServer()
	//service.RegisterUserServiceServer(srv, &Demo{})

	listen, err := net.Listen("tcp", "127.0.0.1:8081")

	if err != nil {
		fmt.Println("listen 8081 failed ")
		return
	}

	err = srv.Serve(listen)

	if err != nil {
		fmt.Println("srv server failed")
		return
	}

	fmt.Println("start ok")
	//s := &service.User{}
	//c := s.GetA()
	//fmt.Println("tt", c)
}
