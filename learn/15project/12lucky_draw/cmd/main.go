package main

import (
	"12lucky_draw/router"
	"12lucky_draw/server"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	handle()
}

func handle() {

	e := gin.New()
	router.InitApi(e)

	ds := server.NewDrawService()
	err := ds.Start()
	if err != nil {
		log.Panic("start draw err :", err)
	}
	log.Println("draw server is starting.................")
	e.Run(":8090")

}
