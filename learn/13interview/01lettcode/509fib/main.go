package main

import "fmt"

func main() {
	n := 5
	res := fib(n)
	fmt.Printf("res=%d\n", res)
}

func fib(n int) int {
	// 递归
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
