package main

import (
	"10go_redis/tcp"
	"log"
)

func main() {
	cfg := &tcp.Config{Address: "0.0.0.0:9090"}
	err := tcp.ListenAndServeWithSignal(cfg, tcp.MakeEchoHandler())
	log.Println("listen err : ", err)
}
