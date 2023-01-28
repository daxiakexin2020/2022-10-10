package service

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	limit := 10
	p := NewPool(3)
	for i := 0; i < limit; i++ {
		//如果没有协程Done，清空管道的一个位置，此处会阻塞
		p.Add(1)
		go func(i int) {
			defer p.Done()
			fmt.Printf("data=%d\n", i)
			time.Sleep(2 * time.Second)
		}(i)
	}
	p.Wait()
	fmt.Println("over")
	fmt.Println("the NumGoroutine continue is:", runtime.NumGoroutine())
}
