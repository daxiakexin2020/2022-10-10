package service

import (
	"fmt"
	"sync"
	"testing"
)

func TestDoOnce(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go doOnce(&wg)
	}
	wg.Wait()
	fmt.Println("main结束")
}

func TestDo(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go do(&wg)
	}
	wg.Wait()
	fmt.Println("main结束")
}
