package main

import (
	"client/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	HandleTest()
}

func HandleTest() {

	cli, err := grpc.Dial("127.0.0.1:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("cli conn failed", err)
	}
	defer cli.Close()

	userClient := service.NewUserServiceClient(cli)

	Req := &service.UserReuqest{
		Id: 1,
	}

	ctx := context.WithValue(context.Background(), "test_ctx_key", "test_ctx_value")
	res, err := userClient.Test(ctx, Req)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Print(res)
}
