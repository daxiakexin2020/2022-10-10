package main

import "07protoc/service/strategys"

func main() {
	test()
}

func test() {
	a := strategys.NewTestA()
	a.Decode("path a")

	b := strategys.NewTestB()
	b.Decode("path b")

	//todo 注册所有策略类
	strategys.RegisterStrategy()

	//todo 从注册树上摘一种策略
	strategy, err := strategys.OFactory.MakeStrategy("testa")
	if err != nil {
		panic(err)
	}
	strategy.Decode("test===========")
}
