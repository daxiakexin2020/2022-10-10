package service

import "sync"

type Pool struct {
	queue chan struct{}
	wg    sync.WaitGroup
}

func NewPool(size uint8) *Pool {
	if size <= 1 {
		size = 1
	}
	return &Pool{
		queue: make(chan struct{}, size),
		wg:    sync.WaitGroup{},
	}
}

func (p *Pool) Add(n int) {
	for j := 0; j < n; j++ {
		//如果没有协程Done，清空管道的一个位置，此处会阻塞
		p.queue <- struct{}{}
	}
	p.wg.Add(n)
}

func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}
