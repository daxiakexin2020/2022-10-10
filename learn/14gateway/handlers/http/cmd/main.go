package main

import (
	"14gateway/components/start_center"
	"14gateway/config"
	"14gateway/db/redis"
	"14gateway/handlers/http/router"
	etcd_service "14gateway/register_center/etcd/service"
	"14gateway/server_govern"
	"fmt"
	"log"
)

func main() {
	if err := initCondition(); err != nil {
		log.Fatalf("initCondition error ： ", err)
	}
	addr := fmt.Sprintf("%s:%d", config.GetWebServerConfig().Addr, config.GetWebServerConfig().Port)
	tmpRegister()
	router.E.Run(addr)
}

func initCondition() error {
	sc := start_center.NewServer()
	sc.Register(redis.InitRedis)
	sc.Register(etcd_service.InitEtcd)
	sc.Register(router.InitApiRouter)
	if err := sc.Run(); err != nil {
		return err
	}
	return nil
}

func tmpRegister() {
	sg := server_govern.NewServer("/test_query", []string{"http://127.0.0.1:9002"}, server_govern.HTTP_TYPE, server_govern.GET_TYPE)
	err := sg.Register()

	sg2 := server_govern.NewServer("/test_json", []string{"http://127.0.0.1:9002"}, server_govern.HTTP_TYPE, server_govern.POST_TYPE)
	err2 := sg2.Register()

	sg3 := server_govern.NewServer("/test_form", []string{"http://127.0.0.1:9002"}, server_govern.HTTP_TYPE, server_govern.POST_TYPE)
	err3 := sg3.Register()
	fmt.Println("临时测试的注册结果2", err, err2, err3)
}
