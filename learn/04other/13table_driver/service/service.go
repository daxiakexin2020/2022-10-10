package service

type Servicer interface {
	Cat() string
	Tail() string
	Set(s string)
}
