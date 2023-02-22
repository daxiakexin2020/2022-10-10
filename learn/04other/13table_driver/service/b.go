package service

type B struct {
	s string
}

func (b *B) Cat() string {
	return "b cat"
}

func (b *B) Tail() string {
	return "b tail"
}

func (b *B) Set(s string) {
	b.s = s
}

func NewB() *B {
	return &B{}
}
