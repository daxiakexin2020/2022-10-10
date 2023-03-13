package main

import (
	"20red_police/config"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
)

func main() {
	run()
}

func run() {
	var wg sync.WaitGroup
	limit := 2
	addr := fmt.Sprintf("%s:%d", config.GetGrpcServerConfig().Addr, config.GetGrpcServerConfig().Port)
	fmt.Println(addr)
	for i := 0; i < limit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dial, err := net.Dial("tcp4", addr)
			if err != nil {
				fmt.Println("dial err:", err)
				return
			}
			_, err = dial.Write([]byte("{\"service_method\":\"Server.Register\",\"meta_data\":{\"name\":\"zz08\",\"pwd\":\"123\",\"repwd\":\"123\",\"phone\":\"45\"}}"))
			if err != nil {
				fmt.Println("dial write err:", err)
			}
			wg.Add(1)
			go read(dial, &wg)
		}()
	}
	wg.Wait()
	fmt.Printf("over,count=%d,successCount=%d,errCount=%d\n", limit, successCount, errCount)
}

var successCount int32
var errCount int32

func read(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
L:
	for {
		b := make([]byte, 1024)
		n, err := conn.Read(b)
		if err != nil && err != io.EOF {
			fmt.Println("read err:", err)
			atomic.AddInt32(&errCount, 1)
			break L
		}
		if n != 0 {
			log.Println("recive server data:", string(b))
			atomic.AddInt32(&successCount, 1)
			break L
		}
	}
}
