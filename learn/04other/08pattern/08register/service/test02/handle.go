package test02

import (
	"errors"
	"fmt"
)

type Test02 struct {
}

func InitTest02() error {
	fmt.Println("test02 start init")
	return errors.New("test02 error")
}
