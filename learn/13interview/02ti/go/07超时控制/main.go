package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	//test01()
	//test02()
	test03()
}

// 等待单个返回值
func test01() {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("i:", i)
			time.Sleep(time.Second * 2)
			ch <- struct{}{}
		}
	}()
	select {
	case <-time.After(time.Second * 1):
		fmt.Println("已经超时")
		return
	case res := <-ch:
		fmt.Println("res:", res)
	}
}

// todo 循环等待，每次等待超时时间重置
func test02() {
	ch := make(chan int)
	go func() {
		for j := 0; j < 10; j++ {
			time.Sleep(time.Second * 1)
			ch <- j
		}
	}()
	for {
		select {
		case <-time.After(time.Second * 4):
			fmt.Println("timeout")
			return
		//todo  进来一次，超时时间，会重置,注意此处的逻辑，因此，会输出0-10.... 然后等待4s之后，超时，4s之内再没有收到数据，则认为超时
		case res := <-ch:
			fmt.Println("data:", res)
		}
	}
}

// context
func test03() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	s := func(ctx context.Context, j int) {
		defer wg.Done()
		select {
		case <-ctx.Done():
			fmt.Println("子任务接收到了超时信号，退出，不再打印了", j)
			return
		default:
			time.Sleep(4 * time.Second)
			fmt.Println("子任务，打印:", j)
		}
	}
	go func() {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			select {
			case <-ctx.Done():
				fmt.Println("主任务接收到了超时信号，退出，不再打印了", j)
				return
			default:
				wg.Add(1)
				go s(ctx, j)
				fmt.Println("主任务，打印:", j)
			}
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Wait()
	fmt.Println("over")
}
