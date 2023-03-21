package user_offline

import (
	"20red_police/asynchronous"
	"io"
	"sync"
)

type UserOffline struct {
	conns map[string]io.WriteCloser
	boot  bool
	mu    sync.Mutex
}

var (
	_  asynchronous.Tasker = (*UserOffline)(nil)
	ec                     = make(chan struct{}, 1)
)

func (uol *UserOffline) Run() error {
	return nil
}

func (uol *UserOffline) Stop() error {
	return nil
}

func (uol *UserOffline) TaskName() string {
	return "USER_OFFLINE"
}

func (uol *UserOffline) ExitSignal() chan struct{} {
	return ec
}
