package main

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
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
	//s := test07(tf01)
	//fmt.Println("S", s)
	//test08()
	//test09()
	test10()
}

func test10() {
	limit := 10
	res := testBackTrack(limit)
	fmt.Printf("最终结果=%d", res)
}

// todo 递归的执行流程和栈一样的，都是后进先出
/**
运动开始时，首先为递归调用建立一个工作栈，其结构包括值参、局部变量和返回地址；

每次执行递归调用之前，把递归函数的值参、局部变量的当前值以及调用后的返回地址压栈；

每次递归调用结束后，将栈顶元素出栈，使相应的值参和局部变量恢复为调用前的值，然后转向返回地址指定的位置继续执行。
*/
func testBackTrack(i int) int {
	fmt.Printf("刚进来的i=%d\n", i)
	if i >= 13 {
		return i
	}
	testBackTrack(i + 1)

	//todo 后进先出
	//递归结束i=12
	//递归结束i=11
	//递归结束i=10
	fmt.Printf("递归结束i=%d\n", i)
	return i
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

func (u *User) ShowName(value interface{}) {
	fmt.Printf("User的名字是=%s\n", u.Name)
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

func test08() {
	res := digui(0)
	fmt.Println(res)
}

func digui(i int) string {
	if i == 5 {
		return "退出了" + strconv.Itoa(i)
	}
	i++
	return digui(i)
}

func test09() {

	data := make(map[int]*User)
	dest := make([]*User, 0)
	data[0] = &User{Name: "zz", Age: 18}
	data[1] = &User{Name: "kx", Age: 16}
	for i := 0; i < 4; i += 2 {
		itme, ok := data[i]
		if ok {
			itme.Name = "被更新了"
		} else {
			itme = &User{"不在", 30}
			dest = append(dest, itme)
		}
	}
	for _, item := range dest {
		fmt.Println("item", item)
	}
	fmt.Println(data)
}
