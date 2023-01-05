package service

import (
	"runtime"
	"testing"
	"time"
)

func TestDoBad(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendBadTasks()
	time.Sleep(time.Second)
	t.Log(runtime.NumGoroutine())
}

func TestDoGood(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	sendGoodTasks()
	time.Sleep(time.Second)
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}

func TestDoCancel(t *testing.T) {
	t.Log(runtime.NumGoroutine())
	doCancel()
	runtime.GC()
	t.Log(runtime.NumGoroutine())
}
