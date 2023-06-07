package main

import (
	"35all_tools/conf"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	app, err := InitApp(gin.New(), conf.GetWebServerConfig())
	if err != nil {
		log.Printf("InitApp err:%v\n", err)
		return
	}
	app.DefaultServerRun()
}
