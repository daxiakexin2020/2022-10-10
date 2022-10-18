package main

import (
	"01option/service"
	"fmt"
	"time"
)

func main() {
	testOption()
}

func testOption() {
	addr := "127.0.0.1"
	port := "6379"
	options := []service.Options{
		service.WithClone(true),
		service.WithDB(1),
		service.WithExpertime(time.Now()),
	}
	r := service.NewClient(addr, port, options...)
	fmt.Println(r)
}
