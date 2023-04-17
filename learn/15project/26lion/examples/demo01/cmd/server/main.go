package main

import (
	"26lion/examples/demo01/apis"
	"26lion/net"
)

func main() {
	run()
}

func run() {
	s := net.NewServer()
	s.Addrouter(1, apis.NewHell())
	s.Serve()
}
