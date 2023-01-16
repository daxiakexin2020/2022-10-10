package main

import "fmt"

/*
*
给定一个整数，写一个函数来判断它是否是 3 的幂次方。如果是，返回 true ；否则，返回 false 。

整数 n 是 3 的幂次方需满足：存在整数 x 使得 n == 3x
*/
func main() {
	n := 9
	res := isPowerOfThree(n)
	fmt.Printf("res=%t", res)
}

func isPowerOfThree(n int) bool {
	if n == 1 || n == 3 {
		return true
	}
	if n <= 0 {
		return false
	}
	if n%3 != 0 {
		return false
	}
	return isPowerOfThree(n / 3)
}
