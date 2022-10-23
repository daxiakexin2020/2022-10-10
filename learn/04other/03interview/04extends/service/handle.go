package service

import "fmt"

type People struct{}

func (p *People) Sleep() {
	fmt.Println("人们睡觉")
}

func (p *People) Say() {
	fmt.Println("人们在说话")
}

type Teacher struct {
	People
}

func NewTeacher() *Teacher {
	return &Teacher{}
}

func (t *Teacher) Say() {
	fmt.Println("老师在说话")
}
