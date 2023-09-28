package main

import (
	"03/server"
	"log"
	"time"
)

func main() {
	test()
}

func test() {

	pool, err := server.GeneratePool(5, 100, 10, 10)
	if err != nil {
		log.Printf("GeneratePool err:%v\n", err)
		return
	}

	defer func() {
		if pool.IsRunning() {
			pool.StopPool()
		}
	}()

	tick := time.Tick(time.Second * 30)

	for i := 0; i < 100; i++ {
		err := pool.AddTask(func() error {
			return nil
		})
		if err != nil {
			log.Printf("add task failed, err:%v\n", err)
		}
	}

L:
	for {
		select {
		case <-tick:
			log.Println("main timeout", pool.AddTaskCount(), pool.ComplateTaskCount(), pool.ActiveCount())
			break L
		default:
			if pool.AddTaskCount() == pool.ComplateTaskCount() {
				log.Println("all task is over", pool.AddTaskCount(), pool.ComplateTaskCount(), pool.ActiveCount())
				break L
			}
		}
	}

	pool.StopPool()
	log.Println("ok............")
}
