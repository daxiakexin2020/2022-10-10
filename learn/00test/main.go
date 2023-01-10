package main

import (
	"fmt"
	"reflect"
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
	//test05()
	//test06()
	s := test07(tf01)
	fmt.Println("S", s)
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

type User struct {
	Name string
	Age  int
}

func test06() {
	//dest := 1
	//dest := []int{1, 2}
	dest := &User{
		Name: "zz",
		Age:  20,
	}
	r := reflect.ValueOf(&dest)
	destValue := reflect.Indirect(r)
	//fmt.Printf("r的值=%+v,destValue的值=%+v\n", r, destValue)
	//fmt.Printf("r的类型=%T,destValue的类型=%T\n", r, destValue)
	fmt.Println(destValue.FieldByName("Name").Interface())
}

type TF func(str string) string

func tf01(str string) string {
	return "tf01" + str
}

func test07(tf TF) string {
	tstr := "abc"
	return tf(tstr)
}
