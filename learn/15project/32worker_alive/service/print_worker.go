package service

import (
	"fmt"
	"time"
)

type PrintWorker struct {
	id  uint
	err error
}

var _ IWorker = (*PrintWorker)(nil)

func NewPrintWorker() *PrintWorker {
	return &PrintWorker{}
}

func (pw *PrintWorker) Work(errc chan<- IWorker) {
	defer func() {
		if r := recover(); r != nil {
			if err, ok := r.(error); ok {
				pw.err = err
			} else {
				pw.err = fmt.Errorf("PRINT WORKER Panic happened with [%v]", r)
			}
		} else {
			pw.err = nil
		}
		errc <- pw
	}()
	fmt.Printf("print worker start work,worker id is :%d \n", pw.id)
	time.Sleep(time.Second * 2)
	panic("worker panic..")
}

func (pw *PrintWorker) ID() uint {
	return pw.id
}

func (pw *PrintWorker) SetID(id uint) {
	pw.id = id
}

func (pw *PrintWorker) Err() error {
	return pw.err
}

func (pw *PrintWorker) SetErr(err error) {
	pw.err = err
}
