package main

import "fmt"

/*
*
todo 如何修改map结构体中的值
*/
func main() {
	//test01()
	//test02()
	test03()
}

func test01() {
	l := &List{
		list: map[string]People{},
	}
	l.list["zz"] = People{Name: "zz", Age: 20}

	//todo 不能修改，是一个值引用，值引用是只读类型，不能修改
	//l.list["zz"].Name = "kx"
	fmt.Println(l)
}

// 方法1
func test02() {
	l := &List{
		list1: map[string]*People{},
	}
	l.list1["zz"] = &People{Name: "zz", Age: 20}

	//todo 可以修改
	l.list1["zz"].Name = "kx"
	fmt.Println(l.list1["zz"])
}

// 方法2
func test03() {
	l := &List{
		list: make(map[string]People),
	}
	l.list["zz"] = People{Name: "zz", Age: 20}

	//todo 可以修改 只是发生了2次拷贝 是先做一次值拷贝，做出一个tmp副本,然后修改该副本，然后再次发生一次值拷贝复制回去
	tmp := l.list["zz"]
	tmp.Name = "kx"

	l.list["zz"] = tmp

	fmt.Println(l)
}

type List struct {
	list  map[string]People
	list1 map[string]*People
}

type People struct {
	Name string
	Age  uint8
}
