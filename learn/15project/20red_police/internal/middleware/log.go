package middleware

import (
	"20red_police/network"
	"fmt"
)

func LogMiddleware(req *network.Request) error {
	fmt.Println("log middler start", req)
	return nil
}
