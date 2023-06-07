package main

import "fmt"

func main() {
	digits := []int{9, 9, 9, 9}
	one := plusOne(digits)
	fmt.Println(one)
}

func plusOne(digits []int) []int {

	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		//找到第一位不等于9的数
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	//每一位都是9
	all9 := make([]int, len(digits)+1)
	all9[0] = 1
	return all9
}
