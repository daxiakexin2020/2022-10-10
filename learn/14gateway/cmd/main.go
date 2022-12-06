package main

import (
	"14gateway/components/start_center"
	redis_service "14gateway/db/redis"
	proxy_http "14gateway/proxy/http"
	etcd_service "14gateway/register_center/etcd/service"
	"fmt"
	"log"
)

func main() {
	test()
}

func test() {

	initCondition()

	reqData := map[string]interface{}{}
	reqData["name"] = "zz_test_name"
	reqData["code"] = "1234"
	reqData["password"] = "abc中"
	pr, _ := proxy_http.NewProxyRequest("http://127.0.0.1:9002/test_json")
	resp, err := pr.POST()
	defer resp.Body.Close()
	defer pr.Close()

	fmt.Println("get err : ", err)

	result, err := proxy_http.Parse(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("result", result)

}

func initCondition() {

	startServer := start_center.NewServer()
	startServer.Register(etcd_service.InitEtcd)
	startServer.Register(redis_service.InitRedis)
	if err := startServer.Run(); err != nil {
		log.Fatalf("*********startServer start failed : %v\n*********", err)
	}
	fmt.Println("********************************initCondition ok********************************")
}
