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
	//token := jwt.NewToken(jwt.TokenTypeAccessToken)
	//type User struct {
	//	Name string `json:"name"`
	//	Age  int    `json:"age"`
	//}
	//u := &User{
	//	Name: "test_zz",
	//	Age:  32,
	//}
	//s, err := token.Make(u)
	//fmt.Println(s)
	//fmt.Println(err)
	//
	//res, _ := jwt.Parse(s, "!")
	//
	//ss := User{}
	//b, er := json.Marshal(res)
	//fmt.Println(er)
	//json.Unmarshal(b, &ss)
	//
	//fmt.Print(ss.Age, ss.Name)

	//fmt.Println("etcd res", config.GetEtcdConfig())
	//fmt.Println("jwt res", config.GetJwtConfig())
	//fmt.Println("mysql res", config.GetMysqlConfig())
	//fmt.Println("redis res", config.GetRedisConfig())
	//fmt.Println("redis res", config.GetRedisConfig())

	//redis test
	//redsiConfig := config.GetRedisConfig()
	//addredss := fmt.Sprintf("%s:%d", redsiConfig.Addr, redsiConfig.Port)
	//rclient, err := redis.NewRedisClient(addredss)
	//fmt.Println("连接redis error", err)
	//fmt.Println("连接redis ", rclient)
	//rclient.Set("test_key_10", "test_value_10", 1*time.Second)
	//v, err := rclient.Get("test_key_10")
	//fmt.Println("获取redis key  err : ", err)
	//fmt.Println("获取redis key  ", v)

	initCondition()

	//s := server_govern.NewServer("tets_server_name", []string{"127.0.0.1:80", "127.0.0.1:81"}, server_govern.HTTP_TYPE, server_govern.ANY_TYPE)
	//err := s.Register()
	//fmt.Println("reigster err : ", err)
	//
	//list, err := s.Discovery()
	//fmt.Println("discovery list err : ", err, list)
	//
	//err = s.UnRegister()
	//fmt.Println("delete  err : ", err)
	//
	//list2, err2 := s.Discovery()
	//fmt.Println("discovery list2 err : ", err2, list2)

	reqData := map[string]interface{}{}
	reqData["name"] = "zz_test_name"
	reqData["code"] = "1234"
	pr, _ := proxy_http.NewProxyRequest("http://127.0.0.1:9002/test_form", reqData)
	pr.SetForm("name", "set")
	pr.SetForm("password", "set")
	pr.SetPostForm("password", "set")

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
