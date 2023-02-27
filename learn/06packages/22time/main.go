package main

import (
	"fmt"
	"time"
)

func main() {

	/**
	todo 这里的 Ticker 跟 Timer 的不同之处，就在于 Ticker 时间达到后不需要人为调用 Reset 方法，会自动续期。
	ticker定时器表示每隔一段时间就执行一次，一般可执行多次。
	timer定时器表示在一段时间后执行，默认情况下只执行一次，如果想再次执行的话，每次都需要调用 time.Reset() 方法，此时效果类似ticker定时器。同时也可以调用 Stop() 方法取消定时器
	timer定时器比ticker定时器多一个 Reset() 方法，两者都有 Stop() 方法，表示停止定时器,底层都调用了stopTimer() 函数。
	除了上面的定时器外，Go 里的 time.Sleep 也起到了类似一次性使用的定时功能。只不过 time.Sleep 使用了系统调用。而像上面的定时器更多的是靠 Go 的调度行为来实现。
	无论哪种计时器，.C 都是一个 chan Time 类型且容量为 1 的单向 Channel，当有超过 1 个数据的时候便会被阻塞，以此保证不会被触发多次。
	*/

	//返回当前时间
	now := time.Now()
	fmt.Println("now:", now)

	//创建一个时间
	unix := time.Unix(1699672271, 0)
	fmt.Println("Unix:", unix)

	//时间类型转为字符串
	s := now.String()
	fmt.Println("s:", s)

	//感觉是垃圾方法....
	parse, err := time.Parse(time.Layout, s)
	fmt.Println("parse:", parse, err)

	//格式化，搞成能看的时间样子，感觉常用
	format := now.Format("2006-01-02 15:04:05")
	fmt.Println("format", format)

	//如果now在time2之前，则返回真，否则返回假
	time2 := time.Now()
	before := now.Before(time2)
	fmt.Println("befoer:", before)

	//如果now在time2之后，则返回真，否则返回假
	after := now.After(time2)
	fmt.Println("after:", after)

	//Add返回时间点now+传入的时间
	add := now.Add(time.Second * 1)
	fmt.Println("add:", add.Format("2006-01-02 15:04:05"))

	//返回now与time2的差值
	sub := time2.Sub(now)
	fmt.Println("sub:", sub.Abs())

	//NewTimer创建一个Timer，它会在最少过去时间段d后到期，向其自身的C字段发送当时的时间,C是一个时间的管道。
	timer := time.NewTimer(time.Second * 2)
	fmt.Println("newTimer:", timer)
	defer timer.Stop()
L:
	for {
		select {
		case ctime := <-timer.C: //2秒后，管道中被发送来数据，不再阻塞
			fmt.Println("ctime:", ctime.Format("2006-01-02 15:04:05"))
			//todo 重启定时器，否则timer执行一次就停了，会阻塞，直接panic，如果用ticker，则会自动续期
			//timer.Reset(time.Second * 1)
			break L
		}
	}

	//timer2 := time.NewTimer(1 * time.Second)
	//for {
	//	timer2.Reset(1 * time.Second) // 这里复用了 timer
	//	select {
	//	case <-timer2.C:
	//		fmt.Println("每隔1秒执行一次")
	//	}
	//}

	//创建新的ticker
	ticker2 := time.NewTicker(time.Second * 1)
	fmt.Println("ticker:", ticker2)

	defer ticker2.Stop()
	for {
		select {
		//todo 会自动续期，和timer的区别，timer需要人为reset()
		case t2 := <-ticker2.C:
			fmt.Println("t2:", t2.Format("2006-01-02 15:04:05"))
		}
	}
}
