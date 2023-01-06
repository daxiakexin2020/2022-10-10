package service

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestNewReader(t *testing.T) {
	str := "a"
	s := strings.NewReader(str)
	r := NewReader(s)
	b := make([]byte, len(str)*2)

	//todo 读取完成以后，清楚buf缓冲区，n2读到了EOF
	n, err := r.Read(b)
	n2, err2 := r.Read(b)
	fmt.Println("data:", n, b, err)    //data: 1 [97 0] <nil>
	fmt.Println("data2:", n2, b, err2) //data2: 0 [97 0] EOF
}

func TestNewWriter(t *testing.T) {

	//创建服务端
	l, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Panic("服务端建立监听失败", err)
	}

	go handleListen(l)

	//连接服务端
	for j := 0; j < 10; j++ {
		tcpw, err := net.DialTimeout("tcp", ":7777", time.Second*10)
		if err != nil {
			log.Panic("tcp连接建立失败\n", err)
		}
		_ = NewWriter(tcpw)
		//b := make([]byte, 1024)
		buf := bytes.Buffer{}
		for k := 0; k < 3; k++ {
			buf.Write([]byte(strconv.Itoa(j)))
		}
		tcpw.Write(buf.Bytes())
	}

	select {}
}

func handleListen(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Println("有连接进来了")
		go read(conn)
	}
}

func read(conn net.Conn) {
	b := make([]byte, 1024)
	n, _ := conn.Read(b)
	t := b[0:n]
	fmt.Println("有连接进来了", n, len(t), string(t))
}

func TestBuffer(t *testing.T) {

	//可以用于字符串拼接
	bb := bytes.Buffer{}
	bb.WriteString("abc")
	bb.WriteString("123")
	fmt.Println(bb.String()) //abc123
}
