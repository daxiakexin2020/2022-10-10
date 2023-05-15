package main

import (
	"32worker_alive/service"
	"log"
)

func main() {
	test01()
}

func test01() {
	manger := service.NewWorkerManger(service.NewTimeWorker(), service.NewPrintWorker(), service.NewTimeWorker())
	if err := manger.StartWorkerPool(); err != nil {
		log.Println("Worker Manage err:", err)
	}
}
