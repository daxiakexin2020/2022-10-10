package generate

import (
	"fmt"
	"testing"
)

type TestStruct struct {
	Name string `zorm:"name"`
	Age  int    `zorm:"age"`
}

func TestB(t *testing.T) {
	data := TestStruct{
		Name: "zz",
	}
	B(&data)
	a := 1
	build, err := B(a)
	fmt.Println("b res : ", build, err)
}
