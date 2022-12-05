package main

import (
	"14gateway/components/start_center"
	redis_service "14gateway/db/redis"
	proxy_http "14gateway/proxy/http"
	etcd_service "14gateway/register_center/etcd/service"
	"fmt"
	"io"
	"log"
)

func main() {
	test()
}

func test() {

	//initCondition()

	reqData := map[string]interface{}{}
	reqData["name"] = "zz_test_name"
	reqData["code"] = "1234"
	pr, _ := proxy_http.NewProxyRequest("http://127.0.0.1:9002/test_json", reqData, proxy_http.JSON_TYPE)
	resp, err := pr.POST()
	fmt.Println("get err : ", err)
	for {
		// 接收服务端信息
		buf := make([]byte, 1024)
		n, err := resp.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println(err)
			return
		} else {
			fmt.Println("读取完毕")
			res := string(buf[:n])
			fmt.Println(res)
			break
		}
	}

	defer resp.Body.Close()
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
