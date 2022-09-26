package main

import (
	"fmt"
	"log"
	"mysql/conf"
	"mysql/driver"
	"mysql/model"
)

func main() {
	handle()
}

func handle() {

	//初始化配置
	conf.Handle()

	//初始化链接
	mc, err := conf.GetMysqlConfig()

	if err != nil {
		log.Fatalf("init mysql config error")
	}

	//链接mysql
	db, err := driver.Handle(mc)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(db)
	model.SetDB(db)

	user := model.User{}
	res := user.Get()
	fmt.Println(res)

	user.Create()

	res = user.Get()
	fmt.Println(res)

	user.Update()

	res = user.Get()
	fmt.Println(res)
}
