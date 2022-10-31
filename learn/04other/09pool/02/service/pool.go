package service

import (
	"fmt"
)

/*******************************task**************************************/

type Task struct {
	handle func() error
}

func NewTask(h func() error) *Task {
	t := &Task{
		handle: h,
	}
	return t
}

func (t *Task) Excetue() {
	t.handle()
}

/*******************************pool**************************************/

type Pool struct {
	cap          int
	JobChannel   chan *Task
	EntryChannel chan *Task
}

func NewPool(cap int) *Pool {
	return &Pool{
		cap:          cap,
		JobChannel:   make(chan *Task),
		EntryChannel: make(chan *Task),
	}
}

func (p *Pool) worker(workId int) {
	for t := range p.JobChannel {
		t.Excetue()
		fmt.Printf("*******************workID=%d执行了任务\n********************", workId)
	}
}

func (p *Pool) run() {
	for i := 0; i < p.cap; i++ {
		go p.worker(i)
	}

	for job := range p.EntryChannel {
		p.JobChannel <- job
	}
}

func (p *Pool) Add(t *Task) {
	for job := range p.EntryChannel {
		p.JobChannel <- job
	}
}

func Test() {
	t := NewTask(func() error {
		fmt.Println("第一个测试协程池的task")
		return nil
	})

	p := NewPool(3)

	go func() {
		for {
			p.EntryChannel <- t
		}
	}()

	p.run()
}
