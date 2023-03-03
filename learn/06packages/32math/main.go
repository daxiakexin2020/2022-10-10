package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	//返回较大指     min/abs
	max := math.Max(6.6, 8.8)
	fmt.Println("max:", max)

	//返回一个有n个元素的，[0,n)范围内整数的伪随机排列的切片。
	perm := rand.Perm(10)
	fmt.Println("perm:", perm)
}
