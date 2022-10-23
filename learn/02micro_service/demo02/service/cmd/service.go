package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata" // grpc metadata包
	"net"
	"service/model"
	"service/service"
)

func main() {
	Handle()
}

var (
	Address = "127.0.0.1"
	Port    = 8080
)

func Handle() {
	addr := fmt.Sprintf("%s:%d", Address, Port)
	fmt.Println(addr)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	//注册拦截器
	var opts []grpc.ServerOption
	opts = append(opts, grpc.UnaryInterceptor(interceptor))
	s := grpc.NewServer(opts...)
	service.RegisterUserServiceServer(s, &model.User{})
	grpclog.Println("Listen on " + Address + " with TLS + Token + Interceptor")
	s.Serve(listen)
}

func auth(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.New("metadata验证失败")
	}

	var (
		appid  string
		appkey string
	)

	if val, ok := md["appid"]; ok {
		appid = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "2012" || appkey != "test key" {
		return errors.New("appid或者appkey验证失败" + appid + appkey)
	}

	return nil
}

//拦截器
func interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	err := auth(ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}
