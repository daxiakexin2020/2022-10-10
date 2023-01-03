package service

import (
	"fmt"
	"sync"
	"time"
)

func doBad(tasks chan int) {
	for {
		select {
		case t := <-tasks:
			time.Sleep(time.Microsecond)
			fmt.Printf("task %d is done\n", t)
		}
	}
}

func sendBadTasks() {
	tasks := make(chan int, 10)
	go doBad(tasks)
	for i := 0; i < 50; i++ {
		tasks <- i
	}
}

func doGood(tasks chan int, wg *sync.WaitGroup) {
	for {
		select {
		case t, beforeClosed := <-tasks:
			if !beforeClosed {
				fmt.Println("taskCh has been closed")
				wg.Done()
				return
			}
			time.Sleep(time.Second)
			fmt.Printf("task %d is done\n", t)
		}
	}
}

func doGoodForRange(tasks chan int, wg *sync.WaitGroup) {
	for t := range tasks {
		time.Sleep(time.Second)
		fmt.Printf("task %d is done\n", t)
	}
	wg.Done()
}

func sendGoodTasks() {
	var wg sync.WaitGroup
	tasks := make(chan int, 10)
	wg.Add(1)
	go doGoodForRange(tasks, &wg)
	for i := 0; i < 20; i++ {
		tasks <- i
	}
	close(tasks)
	fmt.Println("关闭管道")
	wg.Wait()
	fmt.Println("主协程退出")
}
