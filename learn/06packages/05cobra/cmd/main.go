package main

import "05cobra/service"

func main() {
	test()
}

func test() {
	service.Handle()
	service.Run()
}
