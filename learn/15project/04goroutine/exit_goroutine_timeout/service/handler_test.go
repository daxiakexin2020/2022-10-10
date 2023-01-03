package service

import (
	"runtime"
	"testing"
	"time"
)

// go test -run ^TestBadTimeout$ . -v
func TestBadTimeout(t *testing.T) {
	test(t, doBadthing)
}

func Test2phasesTimeout(t *testing.T) {
	for i := 0; i < 1000; i++ {
		timeoutFirstPhase()
	}
	time.Sleep(time.Second * 3)
	t.Log(runtime.NumGoroutine())
}

func test(t *testing.T, f func(chan bool)) {
	t.Helper()
	for i := 0; i < 1000; i++ {
		timeout(f)
	}
	time.Sleep(time.Second * 2)
	t.Log(runtime.NumGoroutine()) //TODO 打印当前协程的个数 1002  即使主协程结束，打印的还是1002个协程
}
