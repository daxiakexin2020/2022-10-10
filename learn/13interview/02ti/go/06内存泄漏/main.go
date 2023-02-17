package main

func main() {

}

func test() {
	/**
	内存泄漏，不是指物理上的内存丢失，而是指由于程序设计不合理，导致部分内存失去控制，回收失败，久而久之，失去控制的内存越来越多，
	可用的内存越来越少，轻则影响性能，重则程序奔溃

	泄漏场景：
		1、time.Ticker造成内存泄漏，间隔一定时间，执行任务，stop一定不能忘记
				func TestTicker(t *testing.T) {
					ticker := time.NewTicker(time.Second)
					todo defer ticker.Stop()	// 这个stop一定不能漏了
					go func(ticker *time.Ticker) {
						for {
							select {
							case value := <-ticker.C:
								fmt.Println(value)
							}
						}
					}(ticker)
					time.Sleep(time.Second * 5)
					fmt.Println("finish!!!")
				}

		2、goroutine配合channel发生泄漏
			例如，使用无缓冲管道
			func TestSend(t *testing.T) {
					ch := make(chan int)
					fmt.Println("num of go start: ", runtime.NumGoroutine())
					time.Sleep(time.Second)

					for i := 0; i < 5; i++ {	// todo 向channel发送5次
						go func(ii int) {
							ch <- ii
							fmt.Println("send to chan: ", ii)
						}(i)
					}

					go func() {		// todo 只从channel接收一次
						value := <-ch
						fmt.Println("recv from chan: ", value)
					}()

					time.Sleep(time.Second)
					fmt.Println("num of go end: ", runtime.NumGoroutine())
			}
			从空的channel接收数据，导致阻塞
			向nil的channel发送或接收
		3、数组的错误使用
			可以使用copy，append 函数，尽量避免，开辟新的内存空间
			尽量在方法内部消化数组/切片，不要传出方法外，如果方法外要使用，可以使用copy，产生新的数组

	监控:
		添加pprof监控代码，定位程序调用堆栈：
	*/
}
