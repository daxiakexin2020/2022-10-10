package main

import "04extends/service"

func main() {
	test()
}

func test() {
	t := service.NewTeacher()
	t.Sleep()
	t.Say()
}
