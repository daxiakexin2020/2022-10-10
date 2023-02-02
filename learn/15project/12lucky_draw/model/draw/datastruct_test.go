package draw

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRand(t *testing.T) {

	fmt.Println("start", time_draw_poll)
	var wg sync.WaitGroup

	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			d, _ := Pick(1)
			if d != 4 {
				fmt.Println(d)
			}
		}()
	}
	wg.Wait()
	fmt.Println("over", time_draw_poll)
	fmt.Println("count", draw_count)
	CountResult()
}

func TestTime(t *testing.T) {
	fmt.Println(time.Now().Hour())
}
