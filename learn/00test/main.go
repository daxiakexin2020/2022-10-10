package main

import (
	"fmt"
	"sort"
)

type Code uint32

type Msg string

func main() {
	//test()
	res := test2()
	fmt.Printf("res=%d", res)
}

func test2() int {
	res := sort.Search(100, func(i int) bool {
		return i >= 20
	})
	return res
}

func test() {
	res := Code(1) //类型转换
	s := Msg("test msg")
	fmt.Println(res, s)
	fmt.Printf("type %T", res)

	t := func(data int32) {
		fmt.Println(data)
	}

	t(int32(res))
}
