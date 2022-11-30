package main

import "etcd/service"

func main() {
	test()
}

func test() {
	key := "test_key_01"
	val := "test_value_01"
	service.TesPut(key, val)
	service.TesWatch(key)
}
