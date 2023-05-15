package service

import (
	"errors"
	"fmt"
)

type WorkerManger struct {
	dieWorkers  chan IWorker
	initWorkers []IWorker
}

func NewWorkerManger(workers ...IWorker) *WorkerManger {
	wm := &WorkerManger{
		dieWorkers:  make(chan IWorker, len(workers)),
		initWorkers: make([]IWorker, 0, len(workers)),
	}
	wm.initWorkers = workers
	return wm
}

func (wm *WorkerManger) StartWorkerPool() error {
	if len(wm.initWorkers) == 0 {
		return errors.New("worker is 0")
	}
	for workId, worker := range wm.initWorkers {
		worker.SetID(uint(workId))
		go worker.Work(wm.dieWorkers)
	}
	wm.keepWorkerAlive()
	return nil
}

func (wm *WorkerManger) keepWorkerAlive() {
	for worker := range wm.dieWorkers {
		fmt.Printf("Worker %d stopped with err: [%v] \n", worker.ID(), worker.Err())
		worker.SetErr(nil)
		go worker.Work(wm.dieWorkers)
	}
}
