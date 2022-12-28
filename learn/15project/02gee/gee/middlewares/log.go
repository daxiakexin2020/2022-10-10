package middlewares

import (
	"gee/gee"
	"log"
	"time"
)

func Log() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		// if a server error occurred
		c.HTML(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
