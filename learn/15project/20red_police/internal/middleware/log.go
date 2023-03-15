package middleware

import (
	"20red_police/network"
	"log"
)

func LogMiddleware(req *network.Request) error {
	log.Println("********************LogMiddleware  req:", req)
	return nil
}
