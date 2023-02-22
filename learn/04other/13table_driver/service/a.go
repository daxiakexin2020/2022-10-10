package service

import "fmt"

type A struct {
	s string
}

func (a *A) Cat() string {
	return a.s + ": a cat"
}

func (a *A) Tail() string {
	return a.s + ": a tail"
}

func (a *A) Set(s string) {
	a.s = s
}

func NewA() *A {
	fmt.Println("a被执行了")
	return &A{}
}
