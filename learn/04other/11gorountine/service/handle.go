package service

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func Handle01() {

	limit := 10
	ch := make(chan int, limit)
	flags := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		go func(k int) {
			ch <- k
			flags <- struct{}{}
		}(i)
	}

	go func() {
		for {
			if len(flags) == limit {
				close(ch)
				break
			}
		}
	}()

	for v := range ch { //关闭的chan，也会一直遍历，直到没有数据，则退出，使用range方式，一定要保证有地方关闭chan，否则会阻塞...
		fmt.Printf("**********data:%d**********\n", v)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("main 结束")

}

func Handle02() {

	limit := 10000
	dataCh := make(chan int, limit)
	flags := make(chan struct{}, limit)
	stop := make(chan struct{}, 1)
	for i := 0; i < limit; i++ {
		go func(k int) {
			dataCh <- k
			flags <- struct{}{}
		}(i)
	}

	go func() {
		for {
			if len(flags) == limit && len(dataCh) == 0 {
				stop <- struct{}{}
				break
			}
		}
	}()

	//单消费者
L:
	for {
		select {
		case data := <-dataCh:
			fmt.Printf("获取到数据data：%d\n", data)
		case <-stop:
			fmt.Println("go协程结束了")
			break L
		default:

		}
	}
	fmt.Println("main结束")
}

func Handle03() {

	startTime := time.Now()
	limit := 100000
	dataCh := make(chan int, limit)
	flags := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		go func(k int) {
			dataCh <- k
			flags <- struct{}{}
		}(i)
	}

	go func() {
		for {
			if len(flags) == limit {
				close(dataCh)
				break
			}
		}
	}()

	var wg sync.WaitGroup
	var total int32

	//多消费者
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			//todo for range channel ,管道关闭后，依然会继续从管道读取数据，一直到读完数据为止
			defer wg.Done()
			for data := range dataCh {
				defer atomic.AddInt32(&total, 1)
				fmt.Printf("data:%d\n", data)
			}
		}()
	}

	wg.Wait()
	fmt.Printf("一共输出%d次\n", total)
	fmt.Println("一共耗时", time.Since(startTime))

	fmt.Println("main结束")
}
