package service

import (
	"fmt"
	"time"
)

func Handle() {

	//定时器，到达一定时间，则发送信号
	timec := time.NewTimer(3 * time.Second)
	select {
	case <-timec.C:
		fmt.Println("到点了")
	}
	fmt.Println("结束")
}

func Handle02() {
	fmt.Println("start", time.Now().Unix())
	<-time.After(1 * time.Second)
	fmt.Println("end", time.Now().Unix())
}

func Handle03() {

	//周期性定时器，每间隔一定时间，将会发送信号到管道，例如定时任务
	timeTicker := time.NewTicker(1 * time.Second)

	defer timeTicker.Stop()

	//每间隔1秒，则会读到定时器发送的信号，如果管道为空，则阻塞
	for range timeTicker.C {
		fmt.Println("now:", time.Now().Unix())
	}
}
