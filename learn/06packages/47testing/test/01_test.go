package test

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	fmt.Println("test01")
}

type People struct {
	Name string
	Age  int
}

func Test02Group(t *testing.T) {
	tests := []People{
		{Name: "zz", Age: 30},
		{Name: "kx", Age: 20},
	}

	//分组，按组批量运行
	for _, p := range tests {
		t.Run(p.Name, func(t *testing.T) {
			fmt.Println("age:", p.Age)
			t.Run("01", Test01)
			t.Run("01", Test01)
		})
	}
}
