package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	//返回v编码后会占用多少字节，注意v必须是定长值、定长值的切片、定长值的指针。
	v := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	size := binary.Size(v)
	fmt.Println("size:", size)
}
