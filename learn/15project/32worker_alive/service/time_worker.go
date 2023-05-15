package service

import (
	"fmt"
	"time"
)

type TimeWorker struct {
	id  uint
	err error
}

var _ IWorker = (*TimeWorker)(nil)

func NewTimeWorker() *TimeWorker {
	return &TimeWorker{}
}

func (tw *TimeWorker) Work(errc chan<- IWorker) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				tw.err = err
			} else {
				tw.err = fmt.Errorf("TIME WORKER Panic happened with [%v]", r)
			}
		} else {
			tw.err = nil
		}
		errc <- tw
	}()
	fmt.Printf("time worker start work,worker id is :%d,time is :%v \n", tw.id, time.Now().Unix())
	time.Sleep(time.Second * 1)
	panic("worker panic..")
}

func (tw *TimeWorker) ID() uint {
	return tw.id
}

func (tw *TimeWorker) SetID(id uint) {
	tw.id = id
}

func (tw *TimeWorker) Err() error {
	return tw.err
}

func (tw *TimeWorker) SetErr(err error) {
	tw.err = err
}
