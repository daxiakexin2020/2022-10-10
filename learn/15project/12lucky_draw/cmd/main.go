package main

import (
	"12lucky_draw/router"
	"12lucky_draw/server"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	start()
}

func start() {
	e := gin.New()
	initRouter(e)
	initDrawPoll()
	log.Println("draw server is starting.................")
	e.Run(":8090")
}

func initRouter(e *gin.Engine) {
	router.InitApi(e)
}

func initDrawPoll() {
	ds := server.NewDrawService()
	if err := ds.Start(); err != nil {
		log.Panic("start draw err :", err)
	}
	ds.SelectTimeDrawPoll()
}
