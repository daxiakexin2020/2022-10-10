package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := true
	sizeof := unsafe.Sizeof(a)
	fmt.Printf("size=%d\n", sizeof)
}
