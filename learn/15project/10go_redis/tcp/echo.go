package tcp

import (
	"10go_redis/lib/sync/atomic"
	"10go_redis/lib/sync/wait"
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"sync"
	"time"
)

type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
}

func MakeEchoHandler() *EchoHandler {
	return &EchoHandler{}
}

type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

func (c *EchoClient) CLose() error {
	c.Waiting.WaitWithTimeout(10 * time.Second)
	c.Conn.Close()
	return nil
}

func (h *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if h.closing.Get() {
		_ = conn.Close()
		return
	}
	client := &EchoClient{Conn: conn}

	//将当前连接存起来
	h.activeConn.Store(client, struct{}{})

	//新建一个io reader
	reader := bufio.NewReader(conn)

	//循环读取发来的信息
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Println("connection close")
				h.activeConn.Delete(client)
			} else {
				log.Println(err)
			}
			return
		}
		client.Waiting.Add(1)
		b := []byte("回复：" + msg)

		//读到的内容，直接吐出去
		_, _ = conn.Write(b)
		client.Waiting.Done()
	}
}

func (h *EchoHandler) Close() error {
	log.Println("shutting down")
	h.closing.Set(true)
	h.activeConn.Range(func(key, value any) bool {
		client := key.(*EchoClient)
		_ = client.CLose()
		return true
	})
	return nil
}
