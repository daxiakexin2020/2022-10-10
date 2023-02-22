package main

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
)

func main() {
	test()
}

func test() {
	e := email.NewEmail()
	e.From = "zz<578975595@qq.com>"
	e.To = []string{"work_docker@aliyun.com"}
	//e.Subject = "Awesome web"   //可选
	e.Text = []byte("hello email Text Body is, of course, supported!")
	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "578975595@qq.com", "oegddlvgapjzbfed", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
}
