package server

import "sync"

type TryLock struct {
	c   chan struct{}
	len int
	mu  sync.Mutex
}

func NewLock(len int) TryLock {
	c := make(chan struct{}, len)
	for i := 0; i < len; i++ {
		c <- struct{}{}
	}
	return TryLock{
		c:   c,
		len: len,
	}
}

func (l *TryLock) Lock() bool {
	select {
	case <-l.c:
		return true
	default:
		return false
	}
}

func (l *TryLock) UnLock() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if len(l.c) == l.len {
		return
	}
	l.c <- struct{}{}
}
