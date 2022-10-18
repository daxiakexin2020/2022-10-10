package main

import "06proxy/service"

func main() {
	handleProxy()
}

func handleProxy() {
	user := &service.User{}
	up := service.NewUserProxy(user)
	username := "test01"
	passwd := "123456"
	up.Login(username, passwd)
}
