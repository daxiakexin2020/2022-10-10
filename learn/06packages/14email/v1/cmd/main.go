package main

import (
	"fmt"
	"v1/server"
)

func main() {
	test()
}

func test() {
	from := "zz<578975595@qq.com>"
	to := []string{"work_docker@aliyun.com"}
	//e.Subject = "Awesome web"   //可选
	text := []byte("hello email Text Body is, of course, supported!")
	e := server.NewEmail(from, to, text, server.WithSubject("zz test"), server.WithHtml([]byte("<span style='color:red;weight:800'>zz test span</span>")))
	sp := server.PlainAuth{
		Identity: "",
		Addr:     "smtp.qq.com:25",
		Username: "578975595@qq.com",
		Pwd:      "oegddlvgapjzbfed",
		Host:     "smtp.qq.com",
	}
	err := e.Send(sp)
	if err != nil {
		fmt.Printf("err=%v", err)
		return
	}
	fmt.Println("ok")
}
