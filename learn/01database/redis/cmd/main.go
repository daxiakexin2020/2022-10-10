package main

import (
	"fmt"
	"log"
	"redis/service"
	"sync"
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

	key := "test_key"

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res, err := client.DispersedLock(key)
			fmt.Println(res, err)
		}()
	}

	wg.Wait()
	fmt.Println("total", service.Total, service.GetCount)
}
