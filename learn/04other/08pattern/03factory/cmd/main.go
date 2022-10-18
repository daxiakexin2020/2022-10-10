package main

import (
	"03factory/service"
	"fmt"
)

func main() {
	testFactory()
}

func testFactory() {

	//ca, err := service.NewArticle(service.C_TYPE)

	fca := service.CreateChine{}
	ca, err := fca.CreateArticle()

	if err != nil {
		panic(err)
	}
	ca.Read()
	ca.Write("测试中文文章01")
	ca.Read()
	ca.Write("测试中文文章02")
	ca.Read()

	//ea, err := service.NewArticle(service.E_TYPE)

	fmt.Println("====================================================")
	fea := service.CreateEngine{}
	ea, err := fea.CreateArticle()
	if err != nil {
		panic(err)
	}
	ea.Read()
	ea.Write("test englise article 01")
	ea.Read()
	ea.Write("test englise article 02")
	ea.Read()
}
