package main

import (
	"fmt"
	"io"
	"strconv"
	"unsafe"
)

type OSConsumer struct {
	p []byte
}

func (osc *OSConsumer) Read(p []byte) (n int, err error) {
	src := []byte("abc")
	//p = append(p, src...)
	copy(p, src)
	fmt.Println("p", p)
	n2, _ := strconv.Atoi(fmt.Sprintf("%d", unsafe.Sizeof(src)))
	return n2, nil
}

func TestReader(reader io.Reader, p []byte) (n int, err error) {
	return reader.Read(p)
}

func main() {

	//io.Reader
	osc := &OSConsumer{}
	p := make([]byte, 10)
	n, err := TestReader(osc, p)
	fmt.Println("io.Reader:", n, p, err)

	//ioutil包中好多函数已经废弃
}
