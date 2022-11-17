package main

import (
	"01/service"
	"01/service/backend"
	"fmt"
	"log"
)

func main() {
	TestServer()
}

func TestServer() {
	s := service.NewServer()
	s.RegisterService(
		backend.NewA(),
		backend.NewB(),
		backend.NewC(),
		backend.NewCon(),
	)
	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("服务启动成功", s.ConH.Port)

	s.AH.SetAge(20)
	fmt.Println("a age", s.AH.ShowAge())
}
