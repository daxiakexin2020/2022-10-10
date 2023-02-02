package tcp

import (
	"10go_redis/interface/tcp"
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Config struct {
	Address    string        `yaml:"address"`
	MaxConnect uint32        `yaml:"max-connect"`
	Timeout    time.Duration `yaml:"timeout"`
}

// 开启服务 优雅关闭 监听退出信号
func ListenAndServeWithSignal(cfg *Config, handler tcp.Handler) error {
	closeCh := make(chan struct{})
	signCh := make(chan os.Signal)

	//监听各种退出信号
	signal.Notify(signCh, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		sign := <-signCh
		switch sign {
		case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			closeCh <- struct{}{}
		}
	}()

	listen, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		return err
	}
	log.Printf("[bind %s=, start listening .....]", cfg.Address)
	ListenAndServe(listen, handler, closeCh)
	return nil
}

func ListenAndServe(listenner net.Listener, handler tcp.Handler, closeCh <-chan struct{}) {

	//监听错误+退出信号，如果select中不设置default，同时没有IO输出的话，会阻塞.... 直到信号到来
	errCh := make(chan error)
	go func() {
		select {
		case <-closeCh:
			log.Println("get exit signle")
		case e := <-errCh:
			log.Printf("accept err:%s", e.Error())
		}
		log.Println("shutting down...")
		_ = listenner.Close()
		_ = handler.Close()
		os.Exit(0)
	}()

	//循环监听连接进来，体现epoll
	var wg sync.WaitGroup
	ctx := context.Background()
	for {
		conn, err := listenner.Accept()
		if err != nil {
			errCh <- err
			break
		}
		log.Println("accept link.......")
		wg.Add(1)
		//有连接进来，开启协程序，处理请求
		go func() {
			defer wg.Done()
			handler.Handle(ctx, conn)
		}()
	}
	wg.Wait()
}
