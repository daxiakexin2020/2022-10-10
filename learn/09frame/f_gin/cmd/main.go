package main

import (
	"f_gin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	initRoute()
}

func initRoute() {
	r := gin.Default()
	router.RegisterRouter(r)
	//监听端口默认为8080
	r.Run(":9002")
}
