生产者 消费者模型

标志生产者结束的生产结束的方式： sync.WaitGroup 或者是使用 make(chan struct,生产数量)
生产结束后，关闭生产管道，此时，不可以写数据，消费者依然可以读数据

标志消费者完成任务的方式： 判断标识的管道长度是否已符合预期
每个消费者读完数据后，向标识结束的管道发送完成任务的标识，一般是struct{}

主协程序阻塞，判断上述消费者中，标识结束的管道长度，是否已经达到预期，如果到达，则退出阻塞，继续后续的流程



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
	//多个消费者
	for j := 0; j < consumerCount; j++ {
		go func(j int) {
			defer func() {
				flag5 <- struct{}{}
				if err := recover(); err != nil {
					fmt.Printf("消费协程模拟异常,err=%v\n", err)
					mainNeedExit = true
				}
			}()
			//如果管道没有关闭，会阻塞等待，接收数据，如果管道关闭了，则会把数据读完，退出阻塞
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

	//更完善，需要加超时和异常处理，否则，如果有消费协程或者是生产协程有异常，则不会达到此条件，会永远阻塞

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