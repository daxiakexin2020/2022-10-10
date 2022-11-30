package main

import (
	"kafka/service/producer"
)

func main() {
	test()
}

func test() {
	producer.Test("zz", "zz_test_msg")
}
