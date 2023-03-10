package main

import (
	"20red_police/config"
	"20red_police/network"
)

func main() {
	run()
}

func run() {

	s := initApp()
	if err := config.InitializeProxyViper(); err != nil {
		panic(err)
	}
	if err := network.Register(s); err != nil {
		panic(err)
	}
	network.Run()
}
