package main

import "fmt"

/*
*
给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4x
*/
func main() {
	n := 12
	res := isPowerOfFour(n)
	fmt.Printf("res=%t", res)
}

func isPowerOfFour(n int) bool {
	if n == 1 || n == 4 {
		return true
	}
	if n <= 0 {
		return false
	}
	if n%4 != 0 {
		return false
	}
	return isPowerOfFour(n / 4)
}
