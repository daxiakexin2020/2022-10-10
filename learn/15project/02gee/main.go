package main

import (
	"gee/gee"
	"gee/gee/middlewares"
	"net/http"
)

func main() {
	g := gee.New()
	g.Use(middlewares.Log())
	g.Use(middlewares.Recovery())
	g.GET("/", index)
	g.GET("/hello/:name", hello)
	api := g.Group("/api")
	{
		api.GET("/hello", hello)
	}
	g.Run(":8888")
}

func index(c *gee.Context) {
	c.HTML(http.StatusOK, "html")
}

func hello(c *gee.Context) {
	c.JSON(http.StatusOK, gee.H{
		"code": 0,
		"msg":  "ok",
		"data": c.Params,
	})
}
