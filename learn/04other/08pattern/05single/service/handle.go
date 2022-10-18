package service

import "sync"

type Single struct {
	Name string
}

var OutSingle *Single

func NewSingle(name string) {
	once := sync.Once{}
	once.Do(func() {
		if OutSingle == nil {
			OutSingle = &Single{
				Name: name,
			}
		}
	})
}
