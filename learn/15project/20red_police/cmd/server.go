package main

import (
	"20red_police/internal/server"
	"20red_police/network"
)

func main() {
	run()
}

func run() {
	s := server.NewServer()
	err := network.Register(s)
	if err != nil {
		panic(err)
	}
	network.Run()
}
