package service

import (
	"fmt"
	"github.com/Jeffail/tunny"
	_ "github.com/Jeffail/tunny"
	"sync"
	"time"
)

/*
*
TODO 每3秒，只会运行3个协程
2023-01-03 14:20:57 k=2
2023-01-03 14:20:57 k=1
2023-01-03 14:20:57 k=0

2023-01-03 14:20:58 k=3
2023-01-03 14:20:58 k=4
2023-01-03 14:20:58 k=5

2023-01-03 14:20:59 k=6
2023-01-03 14:20:59 k=7
2023-01-03 14:20:59 k=8

2023-01-03 14:21:00 k=9
*/
func do() {
	var wg sync.WaitGroup
	flag := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		flag <- struct{}{}
		go func(k int) {
			defer wg.Done()
			fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+" k=%d\n", k)
			time.Sleep(time.Second)
			<-flag
		}(i)
	}
	wg.Wait()
	fmt.Println("main 结束")
}

/*
*
TODO 每3秒，只会运行3个协程
2023-01-03 14:18:10 k=4
2023-01-03 14:18:10 k=9
2023-01-03 14:18:10 k=7

2023-01-03 14:18:11 k=8
2023-01-03 14:18:11 k=5
2023-01-03 14:18:11 k=0

2023-01-03 14:18:12 k=3
2023-01-03 14:18:12 k=1
2023-01-03 14:18:12 k=2

2023-01-03 14:18:13 k=6
*/
func doThirtLib() {
	pool := tunny.NewFunc(3, func(k interface{}) interface{} {
		fmt.Printf(time.Now().Format("2006-01-02 15:04:05")+" k=%d\n", k)
		time.Sleep(time.Second)
		return nil
	})
	defer pool.Close()

	for j := 0; j < 10; j++ {
		go pool.Process(j)
	}
	time.Sleep(time.Second * 4)
}
