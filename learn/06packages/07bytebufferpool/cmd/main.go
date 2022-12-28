package main

import (
	"fmt"
	"github.com/valyala/bytebufferpool"
)

func main() {
	test()
}

func test() {
	b := bytebufferpool.Get()
	b.WriteString("a")
	b.WriteByte(',')
	b.WriteString("b")
	fmt.Println(b.String())
	bytebufferpool.Put(b) //回收了，不再拼接了
	b.WriteString("c")
	fmt.Println(b.String())
	fmt.Printf("type=%T", b.String())
}
