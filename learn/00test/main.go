package main

import "fmt"

type Code uint32

type Msg string

func main() {
	test()
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
