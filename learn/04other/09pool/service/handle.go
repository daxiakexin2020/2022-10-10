package service

import (
	"fmt"
	"sync"
)

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}

func NewPool(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    &sync.WaitGroup{},
	}
}

func (p *Pool) Add(n int) {
	for i := 0; i < n; i++ {
		p.queue <- 1
	}
	fmt.Println("test03")
	fmt.Println("test04")
	fmt.Println("test05")
	p.wg.Add(n)
}

func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}
