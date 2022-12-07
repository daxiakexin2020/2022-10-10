package main

import (
	"14gateway/components/start_center"
	"14gateway/config"
	"14gateway/db/redis"
	"14gateway/handlers/http/router"
	etcd_service "14gateway/register_center/etcd/service"
	"fmt"
	"log"
)

func main() {
	if err := initCondition(); err != nil {
		log.Fatalf("initCondition error ： ", err)
	}
	addr := fmt.Sprintf("%s:%d", config.GetWebServerConfig().Addr, config.GetWebServerConfig().Port)
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
