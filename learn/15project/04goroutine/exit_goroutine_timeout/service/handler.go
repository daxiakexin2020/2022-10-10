package service

import (
	"fmt"
	"time"
)

/**
超时控制在网络编程中是非常常见的，利用 context.WithTimeout 和 time.After 都能够很轻易地实现。
*/

/*
*
最终程序中存在着 1002 个子协程，说明即使是函数执行完成，协程也没有正常退出。
那如果在实际的业务中，我们使用了上述的代码，那越来越多的协程会残留在程序中，最终会导致内存耗尽（每个协程约占 2K 空间），程序崩溃。


TODO 结果：最终程序中存在着 1002 个子协程，说明即使是函数执行完成，协程也没有正常退出。
	那如果在实际的业务中，我们使用了下述的代码，那越来越多的协程会残留在程序中，最终会导致内存耗尽（每个协程约占 2K 空间），程序崩溃。

TODO 仔细分析：done 是一个无缓冲区的 channel，如果没有超时，doBadthing 中会向 done 发送信号，select 中会接收 done 的信号，因此 doBadthing 能够正常退出，子协程也能够正常退出。
	但是，当超时发生时，select 接收到 time.After 的超时信号就返回了，done 没有了接收方(receiver)，
	而 doBadthing 在执行 1s 后向 done 发送信号，由于没有接收者且无缓存区，发送者(sender)会一直阻塞，导致协程不能退出。

TODO：如何避免 设置有缓冲的管道 或者 例如：doGoodthing，使用select向done管道尝试发送消息，如果失败，说明没有了接收方，直接return了
*/

func doBadthing(done chan bool) {
	time.Sleep(time.Second)
	done <- true //todo 此时会阻塞，因为已经没有了接收方，接收方，超时退出了，chan又是没有缓冲的，因此会阻塞
}

func doGoodthing(done chan bool) {
	time.Sleep(time.Second)
	select {
	case done <- true: //TODO 第二种方式，尝试向done发送，如果失败，说明，没有接收方，则return了
	default:
		return
	}
}

// 还有一些更复杂的场景，例如将任务拆分为多段，只检测第一段是否超时，若没有超时，后续任务继续执行，超时则终止。
func do2phases(phase1, done chan bool) {
	time.Sleep(time.Second) //第一段任务
	select {
	case phase1 <- true:
	default:
		return
	}
	fmt.Println("开始第二段任务")
	time.Sleep(time.Second) //第二段
	done <- true
}

/**

这种情况下，就只能够使用 select，而不能能够设置缓冲区的方式了。因为如果给信道 phase1 设置了缓冲区，phase1 <- true 总能执行成功，那么无论是否超时，都会执行到第二阶段，而没有即时返回，这是我们不愿意看到的。对应到上面的业务，就可能发生一种异常情况，向客户端发送了 2 次响应：
任务超时执行，向客户端返回超时，一段时间后，向客户端返回执行结果。
缓冲区不能够区分是否超时了，但是 select 可以（没有接收方，信道发送信号失败，则说明超时了）
*/

func timeoutFirstPhase() error {
	phase1 := make(chan bool)
	done := make(chan bool)
	go do2phases(phase1, done)
	select {
	case <-phase1:
		<-done
		fmt.Println("done")
		return nil
	case <-time.After(time.Millisecond):
		return fmt.Errorf("timeout")
	}
}

func timeout(f func(chan bool)) error {
	done := make(chan bool, 1) //todo 不加缓冲区的话，会阻塞
	go f(done)
	select {
	case <-done:
		fmt.Println("done")
		return nil
	case <-time.After(time.Microsecond):
		fmt.Println("timeout")
		return fmt.Errorf("timeout")
	}
}
