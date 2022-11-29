package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"kafka/service"
)

func main() {
	test()
}

func test() {
	//service.DemoProduct()
	//	service.DemoCunsumer()

	o := service.NewOption(service.WithRequiredAcks(sarama.WaitForAll))
	fmt.Println(o)
}
