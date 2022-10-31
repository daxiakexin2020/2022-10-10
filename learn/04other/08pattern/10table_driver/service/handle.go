package service

import "fmt"

type T1 struct {
	Name string
}

type T2 struct {
	Name string
}

type Ti interface {
	Pname()
	SetName(name string)
}

func (t1 *T1) Pname() {
	fmt.Println("t1 name : " + t1.Name)
}

func (t1 *T1) SetName(name string) {
	t1.Name = name
}

func (t2 *T2) Pname() {
	fmt.Println("t2 name : " + t2.Name)
}
func (t2 *T2) SetName(name string) {
	t2.Name = name
}

func NewT1() *T1 {
	fmt.Println("t1 被实例化了")
	return &T1{}
}

func NewT2() *T2 {
	fmt.Println("t1 被实例化了")
	fmt.Println("t1 main")
	fmt.Println("t1 main2")

	return &T2{}
}

var Tmap = map[int]Ti{
	1: NewT1(),
	2: NewT2(),
}
