package main

import (
	"05single/service"
	"fmt"
	"strconv"
	"time"
)

func main() {
	testSingle()
}

func testSingle() {
	for i := 0; i < 100000; i++ {
		go service.NewSingle(strconv.Itoa(i) + "test")
	}
	time.Sleep(time.Second * 2)
	fmt.Println(service.OutSingle.Name)
	fmt.Println("main over")
}
