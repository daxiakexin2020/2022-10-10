package server

import (
	"errors"
	"log"
	"sync/atomic"
	"time"
)

type Pool struct {
	minCount           int
	maxCount           int
	activeCount        int32
	addTaskCount       int32
	complateTaskCount  int32
	maxWaitWorkerCount int32
	isRunning          bool
	task               chan func() error
	addTime            time.Duration
	errs               chan error
}

var waitingTime = time.Second * 3

var (
	needReduceErr  = errors.New("减少协程")
	paramErr       = errors.New("参数错误")
	poolStoped     = errors.New("Pool已经停止")
	addTaskTimeout = errors.New("加入任务队列超时")
)

func GeneratePool(minCount, maxCount, maxWaitTask, timeout int) (*Pool, error) {
	if maxCount < minCount {
		return nil, paramErr
	}
	p := &Pool{
		minCount:  minCount,
		maxCount:  maxCount,
		task:      make(chan func() error, maxWaitTask),
		isRunning: true,
		addTime:   time.Second * time.Duration(timeout),
	}

	for i := 0; i < p.minCount; i++ {
		p.addWorker()
	}

	go func() {
		for {
			if !p.isRunning {
				break
			}
			p.balance()
		}
	}()

	return p, nil
}

func (p *Pool) balance() {

L:
	for {
		if !p.isRunning {
			break L
		}
		if p.activeCount > int32(p.minCount) && len(p.task) == 0 {
			p.reduceWorker()
		}
		if len(p.task) > 0 && p.activeCount < int32(p.maxCount) {
			p.addWorker()
		}
		time.Sleep(1000)
	}
}

func (p *Pool) addWorker() {
	atomic.AddInt32(&p.activeCount, 1)
	go p.worker()
}

func (p *Pool) reduceWorker() {
	p.task <- func() error {
		return needReduceErr
	}
}

func (p *Pool) worker() {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("worker panic:%v\n", err)
		}
	}()

	time.Sleep(time.Second * 4)

	for {
		t, ok := <-p.task
		if !ok {
			log.Println("worker out...........................")
			break
		}
		err := t()
		if err != needReduceErr {
			atomic.AddInt32(&p.complateTaskCount, 1)
		}
		if err == needReduceErr && p.activeCount > int32(p.minCount) {
			atomic.AddInt32(&p.activeCount, -1)
			break
		}
	}
}

func (p *Pool) StopPool() {
	p.isRunning = false
	close(p.task)
	time.Sleep(waitingTime) //等待资源回收
}

func (p *Pool) AddTask(t func() error) error {
	if !p.isRunning {
		return poolStoped
	}
	select {
	case p.task <- t:
		atomic.AddInt32(&p.addTaskCount, 1)
		return nil
	case <-time.After(p.addTime):
		return addTaskTimeout
	}
}

func (p *Pool) ComplateTaskCount() int32 {
	return p.complateTaskCount
}

func (p *Pool) AddTaskCount() int32 {
	return p.addTaskCount
}

func (p *Pool) ActiveCount() int32 {
	return p.activeCount
}

func (p *Pool) IsRunning() bool {
	return p.isRunning
}
