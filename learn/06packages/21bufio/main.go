package main

import (
	"bufio"
	"fmt"
	"strconv"
	"unsafe"
)

type OSConsumer struct {
	//p []byte
}

func (osc *OSConsumer) Read(p []byte) (n int, err error) {
	src := []byte("abc")
	p = append(p, src...)
	fmt.Println("p", p)
	n2, _ := strconv.Atoi(fmt.Sprintf("%d", unsafe.Sizeof(src)))
	return n2, nil
}

func main() {
	/**
	bufio包实现了有缓冲的I/O。它包装一个io.Reader或io.Writer接口对象，创建另一个class也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。
	*/

	//创建新的reader包装对象
	p := make([]byte, 10)
	newReader := bufio.NewReader(&OSConsumer{})
	n, err := newReader.Read(p)
	fmt.Println("newReader read:", n, p, err)

	//创建新的scanner
	//scanner := bufio.NewScanner(&OSConsumer{})
	//scan := scanner.Scan()
	//fmt.Println("scanner Scan:", scan)
}
