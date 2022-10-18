package main

import (
	"07facade/service"
	"fmt"
)

func main() {
	handleFacade()
}

func handleFacade() {
	phone := "13681576396"
	code := 1234
	us := service.NewUserService()
	//u1, err1 := us.Register(phone, code)
	//fmt.Println(u1, err1)
	//u2, err2 := us.Login(phone, code)
	//fmt.Println(u2, err2)
	u3, err3 := us.LoginOrRegister(phone, code)
	fmt.Println(u3, err3)
}
