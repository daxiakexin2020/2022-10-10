package main

import (
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

type People struct {
	Name  string
	Age   int
	hobby []string
	//Mutex是一个互斥锁，可以创建为其他结构体的字段；零值为解锁状态。Mutex类型的锁和线程无关，可以由不同的线程加锁和解锁。
	mu sync.Mutex
	//RWMutex是读写互斥锁。该锁可以被同时多个读取者持有或唯一个写入者持有。RWMutex可以创建为其他结构体的字段；零值为解锁状态。RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。
	rwmu sync.RWMutex
}

func main() {

	//once
	var tonce sync.Once
	var wg sync.WaitGroup
	limit := 1 << 10
	var count int64
	for i := 0; i < limit; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			defer atomic.AddInt64(&count, 1)
			tonce.Do(func() {
				fmt.Println("go:" + strconv.Itoa(j)) //只输出1次
			})
		}(i)
	}
	wg.Wait()
	fmt.Printf("once over count=%d\n", count)

	//pool
	p := sync.Pool{
		New: func() interface{} {
			return &People{}
		},
	}

	get := p.Get()
	if pe, ok := get.(*People); ok {
		pe.Name = "Zz"
		fmt.Printf("get1=%+v\n", pe)
		p.Put(get)
	} else {
		fmt.Println("pe failed")
	}

	get2 := p.Get()
	if pe2, ok := get2.(*People); ok {
		pe2.Name = "KX"
		fmt.Printf("get2=%+v\n", pe2)
		p.Put(get2)
	} else {
		fmt.Println("pe2 failed")
	}

	//mu
	pe3 := p.Get().(*People)
	var wg3 sync.WaitGroup
	var count3 int64
	for i := 0; i < 1<<10; i++ {
		wg3.Add(1)
		go func() {
			pe3.mu.Lock()
			defer wg3.Done()
			defer atomic.AddInt64(&count3, 1)
			defer pe3.mu.Unlock()
			pe3.hobby = append(pe3.hobby, strconv.Itoa(i))
		}()
	}
	wg3.Wait()
	//p.Put(pe3)
	fmt.Printf("mu over count3=%d,len=%d\n", count3, len(pe3.hobby))

	//channel
	pe4 := p.Get().(*People)
	var wg4 sync.WaitGroup
	var pcount4 int64
	var ccount4 int64
	data4 := make(chan string, 1<<10)
	flag4 := make(chan struct{}, 1)

	go func() {
		//如果管道没有关闭，会阻塞等待，接收数据，如果管道关闭了，则会把数据读完，退出阻塞
		defer func() {
			flag4 <- struct{}{}
		}()
		for res := range data4 {
			atomic.AddInt64(&ccount4, 1)
			pe4.hobby = append(pe4.hobby, res)
		}
	}()

	for i := 0; i < 1<<10; i++ {
		wg4.Add(1)
		go func() {
			defer wg4.Done()
			defer atomic.AddInt64(&pcount4, 1)
			data4 <- strconv.Itoa(i)
		}()
	}
	wg4.Wait()
	close(data4)
	for len(flag4) < 1 {
	}

	//p.Put(pe4)
	fmt.Printf("mu over pcount4=%d,ccount4=%d,len=%d\n", pcount4, ccount4, len(pe4.hobby))

	//channel 多消费者
	timer := time.NewTimer(time.Second * 3)

	pe5 := p.Get().(*People)
	var wg5 sync.WaitGroup
	var pcount5 int64
	var ccount5 int64
	var mainNeedExit bool
	data5 := make(chan string, 1<<10)
	consumerCount := 1 << 5
	flag5 := make(chan struct{}, consumerCount)
	for j := 0; j < consumerCount; j++ {
		go func(j int) {
			defer func() {
				flag5 <- struct{}{}
				if err := recover(); err != nil {
					fmt.Printf("消费协程模拟异常,err=%v\n", err)
					mainNeedExit = true
				}
			}()
			for res := range data5 {
				if j == 10 {
					time.Sleep(time.Second * 4)
				}
				if j%5 == 0 {
					panic("第" + strconv.Itoa(j) + "消费协程模拟异常:")
				}
				atomic.AddInt64(&ccount5, 1)
				pe5.mu.Lock()
				pe5.hobby = append(pe5.hobby, res)
				pe5.mu.Unlock()
			}
		}(j)
	}

	for i := 0; i < 1<<10; i++ {
		wg5.Add(1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Printf("生产任务异常err=%v\n", err)
					mainNeedExit = true
				}
				wg5.Done()
				atomic.AddInt64(&pcount5, 1)
			}()
			data5 <- strconv.Itoa(i)
		}()
	}
	wg5.Wait()
	close(data5)

L:
	for {
		select {
		case <-timer.C:
			fmt.Println("任务已经超时,超时退出")
			break L
		default:
			if len(flag5) == consumerCount {
				fmt.Println("任务已经完成，正常退出")
				break L
			}
			if mainNeedExit {
				fmt.Println("任务有异常，异常退出")
				break L
			}
		}
	}

	p.Put(pe5)

	fmt.Printf("mu over pcount5=%d,ccount5=%d,len=%d\n", pcount5, ccount5, len(pe5.hobby))

}
