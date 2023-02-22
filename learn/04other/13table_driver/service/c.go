package service

type C struct {
	s string
}

func (c *C) Cat() string {
	return "c cat"
}

func (c *C) Tail() string {
	return "c tail"
}

func (c *C) Set(s string) {
	c.s = s
}

func NewC() *C {
	return &C{}
}
