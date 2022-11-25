package main

import (
	"manage_center"
	"user_center/http"
	"user_center/router"
)

func main() {
	start()
}

func start() {
	router.BootRouter()
	m := manage_center.NewManager()
	m.RegisterKernel(http.Initkernel())
	m.Run()
}
