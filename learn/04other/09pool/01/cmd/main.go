package main

import (
	"01/service"
	"fmt"
	"time"
)

func main() {
	Handle()
}

func Handle() {
	pool := service.NewPool(5)
	limit := 20
	for i := 0; i < limit; i++ {
		pool.Add(1)
		go func(n int) {
			fmt.Printf("i=%d\n", n)
			time.Sleep(time.Second * 1)
			//fmt.Println("the NumGoroutine continue is:", runtime.NumGoroutine(), n)
			pool.Done()
		}(i)
	}
	pool.Wait()
	fmt.Println("====================ok====================")
}
