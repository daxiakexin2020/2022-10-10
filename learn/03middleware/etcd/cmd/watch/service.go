package main

import "etcd/service"

func main() {
	test()
}

func test() {
	key := "test_key_01"
	service.TesWatch(key)
}
