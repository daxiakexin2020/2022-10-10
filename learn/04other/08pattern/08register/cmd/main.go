package main

import (
	"08register/service"
	"08register/service/test01"
	"08register/service/test02"
	"fmt"
	"log"
)

func main() {
	handle()
}

func handle() {

	register := service.GetProxyRegiseter()
	register.Register(test01.InitTest01, test02.InitTest02)
	register.Run()
	err := register.CheckErr()
	if err != nil {
		log.Fatalf("err=%v", err)
	}

	fmt.Println("***************************init all ok***************************")
}
