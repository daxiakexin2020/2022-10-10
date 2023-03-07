package server

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"testing"
)

func TestTryLock(t *testing.T) {

	tests := []int{1, 2, 3}
	for _, test := range tests {
		t.Run(strconv.Itoa(test), func(t *testing.T) {
			var count int32
			var wg sync.WaitGroup
			fmt.Printf("开启%d把锁\n", test)
			l := NewLock(test)
			for i := 0; i < test*1000; i++ {
				wg.Add(1)
				go func() {
					lock := l.Lock()
					defer wg.Done()
					if lock {
						atomic.AddInt32(&count, 1)
						l.UnLock()
					}
				}()
			}
			wg.Wait()
			fmt.Printf("over=%d,有%d个协程拿到锁\n", test, count)
			fmt.Println("****************************************")
		})
	}
}
