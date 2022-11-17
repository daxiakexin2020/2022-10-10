package main

import (
	"03/center"
	"03/center/defined"
	"fmt"
)

func main() {
	test()
}

func test() {

	var ltype center.Ltype = "telephone"
	pl := center.NewProxyLogin(ltype)
	req := &defined.LoginRequest{}
	err := pl.Login(req)
	fmt.Println("登陆结果：", err)

	var ltype2 center.Ltype = "username"
	pl2 := center.NewProxyLogin(ltype2)
	req2 := &defined.LoginRequest{}
	err2 := pl2.Login(req2)
	fmt.Println("登陆结果：", err2)

	var ltype3 center.Ltype = "username"

	pl2.ChengeLtype(ltype3)
	req3 := &defined.LoginRequest{}
	err3 := pl2.Login(req3)
	fmt.Println("登陆结果：", err3)

}
