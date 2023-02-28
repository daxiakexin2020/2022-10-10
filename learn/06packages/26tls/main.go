package main

import (
	"crypto/tls"
	"fmt"
)

func main() {

	//建立连接
	dial, err := tls.Dial("tcp", "www.baidu.com:443", &tls.Config{})
	fmt.Println("tls.Dial:", dial, err)

	//获取一个client
	client := tls.Client(dial.NetConn(), &tls.Config{})
	fmt.Println("tls.Client:", client)

	//RemoteAddr返回远端网络地址。
	s := client.RemoteAddr().String()
	fmt.Println("remoteAddr:", s)

	//LocalAddr返回本地网络地址。
	s2 := client.LocalAddr().String()
	fmt.Println("localAddr:", s2)

	b := make([]byte, 1024)
	client.Read(b)
	fmt.Println("client read:", b, string(b))

	//tls实现细节
	state := client.ConnectionState()
	fmt.Println("client.ConnectionState:", state, state.ServerName, state.Version)
}
