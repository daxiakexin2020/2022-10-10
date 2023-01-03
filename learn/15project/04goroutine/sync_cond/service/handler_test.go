package service

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	cond := sync.NewCond(&sync.Mutex{})
	go read("reader1", cond)
	go read("reader2", cond)
	go read("reader3", cond)
	write("writer", cond)
	time.Sleep(time.Second * 4)
	fmt.Println("main结束")
}
