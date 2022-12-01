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
	s, err := service.NewRegisteServer([]string{"127.0.0.1:2379"}, service.ETCD)
	if err != nil {
		log.Fatalf("获取服务实体错误=>>>%v", err)
	}
	if err = s.Register("test_key", []string{"127.0.0.1:81"}); err != nil {
		log.Fatalf("注册服务失败=>>>", err)
	}
	fmt.Println("服务注册成功")

	v, err := s.Get("test_key")
	if err != nil {
		log.Fatalf("获取val失败", err)
	}

	fmt.Println("获取到的数据", v)
}
