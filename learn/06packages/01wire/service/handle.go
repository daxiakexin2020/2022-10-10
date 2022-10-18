package service

import "fmt"

type A struct {
	Msg string
}

type B struct {
	AY A
}

type C struct {
	BY B
}

func NewA(msg string) *A {
	return &A{
		Msg: msg,
	}
}

func NewB(a A) *B {
	return &B{
		AY: a,
	}
}

func NewC(b B) *C {
	return &C{
		BY: b,
	}
}

func (b *B) Show() {
	fmt.Println(b.AY.Msg)
}

func (c *C) Show() {
	c.BY.Show()
}
