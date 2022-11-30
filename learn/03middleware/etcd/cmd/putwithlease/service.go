package main

import "etcd/service"

func main() {
	test()
}

func test() {
	key := "test_key_02"
	val := "test_value_02"
	service.TesPutWithLease(key, val)
}
