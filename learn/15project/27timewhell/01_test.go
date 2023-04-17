package _7timewhell

import (
	"fmt"
	"testing"
)

func Test01(t *testing.T) {
	step := 2
	index := -1
	num := 0
	circle := 60
	for i := 0; i < 62; i++ {
		index++
		v := index % step
		fmt.Println("v:", index, step, v)
		num = index * step
		fmt.Println("num:", num)
		n := index % circle
		fmt.Println("n:", n)
	}
}
