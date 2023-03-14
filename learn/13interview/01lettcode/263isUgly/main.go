package main

import "fmt"

func main() {
	num := 6
	ugly := isUgly(num)
	fmt.Printf("res=%t", ugly)
}

var factors = []int{2, 3, 5}

func isUgly(n int) bool {
	if n <= 0 {
		return false
	}
	for _, f := range factors {
		for n%f == 0 {
			n /= f
		}
	}
	return n == 1
}
