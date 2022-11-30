package main

import "etcd/service"

func main() {
	test()
}

func test() {
	key := "test_key_02"
	service.TesGet(key)
}
