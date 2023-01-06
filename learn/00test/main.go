package main

import (
	"fmt"
	"sort"
	"time"
)

type Code uint32

type Msg string

func main() {

	//todo-gcflags=-m参数 查看具体堆栈情况
	//test()
	//res := test2()
	//fmt.Printf("res=%d", res)
	//res := test03()
	//res := test04()
	//fmt.Printf("res=%+v,%p", res, res)
	test05()
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

func test03() []int {
	a := []int{1, 2, 3, 4}
	b := make([]int, len(a))
	c := a[0:4]
	copy(b, a)
	_ = c
	//fmt.Printf("%p %p %p \n", a, b, c)
	//fmt.Println(a, b, c)
	//fmt.Println(&a[0], &b[0], &c[0])
	return b
}

type C struct {
	Name string
}

func test04() *C {
	cccccccccccccc := &C{
		Name: "test",
	}
	return cccccccccccccc
}

func test05() {
	d := time.Second * 2
	for {
		t := time.NewTimer(d)
		<-t.C
		fmt.Println("time:", time.Now().Format("2006-01-02 15:04:05"))
	}
}
