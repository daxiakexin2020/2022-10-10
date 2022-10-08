package main

import (
	"client/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"time"
)

func main() {
	Handle()
}

var (
	Addres = "127.0.0.1"
	Port   = 8080
)

// customCredential 自定义认证
type customCredential struct{}

// GetRequestMetadata 实现自定义认证接口
func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "2012",
		"appkey": "test key",
	}, nil
}

// RequireTransportSecurity 自定义认证是否开启TLS
func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func Handle() {

	var opts []grpc.DialOption

	// 指定自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))
	// 指定客户端interceptor
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))

	opts = append(opts, grpc.WithInsecure())

	addr := fmt.Sprintf("%s:%d", Addres, Port)

	dial, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	defer dial.Close()

	c := service.NewUserServiceClient(dial)
	req := &service.UserRequest{
		Id: 1,
	}

	info, err := c.GetUserInfo(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(info)
}

// interceptor 客户端拦截器
func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)
	grpclog.Printf("method=%s req=%v rep=%v duration=%s error=%v\n", method, req, reply, time.Since(start), err)
	return err
}
