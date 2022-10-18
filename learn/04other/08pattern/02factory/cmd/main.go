package main

import "02factory/service"

func main() {
	testFactory()
}

func testFactory() {

	ca, err := service.NewArticle(service.C_TYPE)
	if err != nil {
		panic(err)
	}
	ca.Read()
	ca.Write("测试中文文章01")
	ca.Read()
	ca.Write("测试中文文章02")
	ca.Read()

	ea, err := service.NewArticle(service.E_TYPE)
	if err != nil {
		panic(err)
	}
	ea.Read()
	ea.Write("test englise article 01")
	ea.Read()
	ea.Write("test englise article 02")
	ea.Read()
}
