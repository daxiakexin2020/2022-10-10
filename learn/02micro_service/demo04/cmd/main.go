package main

import (
	"demo04/service"
	"fmt"
	"log"
)

func main() {
	test()
}

func test() {
	s, err := service.NewService("test_service", []string{"127.0.0.1:80"}, service.ETCD)
	if err != nil {
		log.Fatalf("获取服务实体错误=>>>", err)
	}
	if err = s.Register(); err != nil {
		log.Fatalf("注册服务失败=>>>", err)
	}
	fmt.Println("服务注册成功")

	key, _ := s.Gerenral()
	v, err := s.Get(key)
	if err != nil {
		log.Fatalf("获取val失败", err)
	}

	fmt.Println("获取到的数据", v)
}
