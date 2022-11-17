package backend

import "fmt"

type A struct {
	Age int
}

func NewA() *A {
	return &A{}
}

func (a *A) Name() string {
	return "ah"
}

func (a *A) InitService() error {
	fmt.Println("*******************AH服务	init	ok******************")
	a.Age = 10
	return nil
}

func (a *A) ShowAge() int {
	return a.Age
}

func (a *A) SetAge(age int) {
	a.Age = age
}
