package backend

import "fmt"

type C struct {
}

func NewC() *C {
	return &C{}
}

func (c *C) Name() string {
	return "CH"
}

func (c *C) InitService() error {
	fmt.Println("*******************CH服务	init	ok******************")
	return nil
}
