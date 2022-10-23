package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"sync"
)

func main() {
	//Handle()
	HandleChannel()
}

type TestMutex struct {
	data map[string]string
	m    *sync.RWMutex
}

func Handle() {

	tm := &TestMutex{
		data: make(map[string]string),
		m:    new(sync.RWMutex),
	}

	limit := 100
	flag := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		go func() {
			tm.m.Lock()
			defer func() {
				flag <- struct{}{}
			}()
			key := strconv.Itoa(i)
			b := bytes.Buffer{}
			b.WriteString("test_")
			b.WriteString(key)
			b.WriteString("\n")
			tm.data[key] = b.String()
			tm.m.Unlock()
		}()
	}

	for {
		if len(flag) == limit {
			break
		}
	}
	fmt.Println("main over")
	fmt.Println(tm.data, len(tm.data), len(flag))
}

func HandleChannel() {
	limit := 100000
	res := make(chan int, limit)
	flag := make(chan struct{}, 1)
	tm := &TestMutex{
		data: make(map[string]string),
	}
	go func() {
		for {
			data, ok := <-res
			//TODO 生产者做完任务了，消费者读取完数据，发送完成信号至通道，主协程进行退出
			if !ok && data == 0 {
				flag <- struct{}{}
				break
			}
			s := strconv.Itoa(data)
			b := bytes.Buffer{}
			b.WriteString("test_")
			b.WriteString(s)
			tm.data[s] = b.String()
		}
	}()

	wg := sync.WaitGroup{}
	for i := 1; i <= limit; i++ {
		wg.Add(1)
		go func(data int) {
			defer wg.Done()
			res <- data
		}(i)
	}
	wg.Wait()
	close(res)

	//消费者完成任务了，主协程退出
	for {
		if len(flag) == 1 {
			break
		}
	}

	fmt.Println("main over")
	fileName := "F:/go_project/2022-10-10/learn/04other/03interview/02go/03mutex/file/text"
	file, _ := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0777)
	for k, _ := range tm.data {
		io.WriteString(file, k+"\n")
	}
}
