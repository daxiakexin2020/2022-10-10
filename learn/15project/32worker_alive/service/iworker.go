package service

type IWorker interface {
	Work(c chan<- IWorker)
	ID() uint
	SetID(id uint)
	SetErr(err error)
	Err() error
}
