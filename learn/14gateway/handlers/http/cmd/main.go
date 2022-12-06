package main

import (
	"14gateway/components/start_center"
	"14gateway/handlers/http/router"
	"log"
)

func main() {
	if err := initCondition(); err != nil {
		log.Fatalf("initCondition error ： ", err)
	}
	router.E.Run(":6666")
}

func initCondition() error {
	sc := start_center.NewServer()
	sc.Register(router.InitApiRouter)
	if err := sc.Run(); err != nil {
		return err
	}
	return nil
}
