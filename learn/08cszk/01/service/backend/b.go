package backend

import (
	"errors"
	"fmt"
)

type B struct {
	Notify string
}

func NewB() *B {
	return &B{}
}

func (b *B) Name() string {
	return "BH"
}

func (b *B) InitService() error {
	b.Notify = "b的初始化"
	fmt.Println("*******************BH服务	init	ok******************")
	return nil
	return errors.New("B启动失败")
}
