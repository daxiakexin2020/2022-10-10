package main

import (
	"fmt"
	"net"
)

func main() {
	if err := client(); err != nil {
		panic(err)
	}
}

func client() error {
	dial, err := net.Dial("tcp", ":9110")
	if err != nil {
		return err
	}
	dial.Write([]byte("hello tcp"))
	for {
		b := make([]byte, 100)
		dial.Read(b)
		fmt.Println("client read：", string(b))
	}
	return nil
}
