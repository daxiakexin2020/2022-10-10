package main

import (
	"04strategy/service"
	"fmt"
)

func main() {
	testStrategy()
}

func testStrategy() {
	s := &service.S2{}
	ow := service.NewOW(s)
	var m float64 = 200
	hmoney := ow.Handle(m)
	fmt.Printf("优惠后的单价是：%v", hmoney)
}
