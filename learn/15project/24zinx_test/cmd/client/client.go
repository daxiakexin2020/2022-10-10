package main

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"os"
	"os/signal"
	"time"
)

// 创建连接的时候执行
func DoClientConnectedBegin(conn ziface.IConnection) {
	zlog.Debug("DoConnecionBegin is Called ... ")

	//设置两个链接属性，在连接创建之后
	conn.SetProperty("Name", "刘丹冰")
	conn.SetProperty("Home", "https://yuque.com/aceld")

	go business(conn)
}

// 客户端自定义业务
func business(conn ziface.IConnection) {
	for {
		err := conn.SendMsg(100, []byte("zz Ping...[FromClient]"))
		if err != nil {
			fmt.Println(err)
			zlog.Error(err)
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	client := znet.NewClient("127.0.0.1", 9999)
	client.SetOnConnStart(DoClientConnectedBegin)
	client.Start()
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Println("===exit===", sig)
	// 清理客户端
	client.Stop()
	time.Sleep(time.Second * 2)
}
