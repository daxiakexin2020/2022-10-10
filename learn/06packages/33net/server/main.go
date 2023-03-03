package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {

	//创建监听
	listen, err := net.Listen("tcp", ":9110")
	if err != nil {
		return err
	}
	//等待客户端连接
	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		//处理连接
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	reader := bufio.NewReader(conn)
	b := make([]byte, 1024)
	n, err := reader.Read(b)
	fmt.Println("raed:", n, err, string(b))
	conn.Write([]byte("server send hello"))
}
