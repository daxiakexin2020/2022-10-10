package service

import "fmt"

type People struct {
	Name string
	Age  int
	C    *colorManage //享元
}

func NewPeople(name string, age int) *People {
	return &People{
		Name: name,
		Age:  age,
		C:    GetColorManage(),
	}
}

func (p *People) Eat() {
	fmt.Printf("***********************皮肤是%s的%s在吃饭***********************\n", p.C.color, p.Name)
}
