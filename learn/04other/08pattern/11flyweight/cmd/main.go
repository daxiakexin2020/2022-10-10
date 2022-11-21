package main

import (
	"11flyweight/service"
	"fmt"
)

func main() {
	test()
}

func test() {
	var name = "zz"
	var age = 20
	p := service.Handle(name, age)
	p.Eat()

	var name2 = "kx"
	var age2 = 19
	p2 := service.Handle(name2, age2)
	p2.Eat()

	fmt.Println("地址信息", &p.C, &p2.C)

	if p.C == p2.C {
		fmt.Println("享元成功，是同一个对象")
	} else {
		fmt.Println("享元失败，不是同一个对象", &p.C, &p2.C)
	}
}
