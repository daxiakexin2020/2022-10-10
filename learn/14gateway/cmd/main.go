package main

import (
	"14gateway/config"
	"14gateway/db/redis"
	"fmt"
	"time"
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

	fmt.Println("etcd res", config.GetEtcdConfig())
	fmt.Println("jwt res", config.GetJwtConfig())
	fmt.Println("mysql res", config.GetMysqlConfig())
	fmt.Println("redis res", config.GetRedisConfig())
	fmt.Println("redis res", config.GetRedisConfig())

	//redis test
	redsiConfig := config.GetRedisConfig()
	addredss := fmt.Sprintf("%s:%d", redsiConfig.Addr, redsiConfig.Port)
	rclient, err := redis.NewRedisClient(addredss)
	fmt.Println("连接redis error", err)
	fmt.Println("连接redis ", rclient)
	rclient.Set("test_key_10", "test_value_10", 1*time.Second)
	v, err := rclient.Get("test_key_1011")
	fmt.Println("获取redis key  err : ", err)
	fmt.Println("获取redis key  ", v)

}
