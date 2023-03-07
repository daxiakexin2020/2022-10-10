package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	//原子内存级别的操作
	var wg sync.WaitGroup
	var count int64
	limit := 100000
	for i := 0; i < limit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&count, 1)
		}()
	}
	wg.Wait()
	fmt.Printf("over,count=%d\n", count)
}
