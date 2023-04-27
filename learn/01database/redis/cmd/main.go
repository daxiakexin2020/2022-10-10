package main

import (
	"fmt"
	"log"
	"redis/bus"
	"redis/service"
)

func main() {
	test()
}

func test() {

	addr := "127.0.0.1:6379"

	//业务应该保存此clinet，去使用
	client, err := service.NewClient(addr)

	if err != nil {
		log.Fatalf("连接redis失败->>>>>%v", err)
	}

	id := "val_id"
	name := "val_name"
	token := "val_token"
	auth := bus.NewAuth(id, name, token, bus.WithNamesapce("val_namespce"))
	err = auth.Set(client)
	fmt.Println("set auth err:", err)

	get, err := client.HGet(auth.Token, bus.Auth_Namespace)
	fmt.Println("HGet:::::", get, err)

	del, err := client.Del(auth.Token)
	fmt.Println("Del:::::", del, err)

	get2, err2 := client.HGet(auth.Token, bus.Auth_Name)
	fmt.Println("HGet2:::::", get2, err2)

}
