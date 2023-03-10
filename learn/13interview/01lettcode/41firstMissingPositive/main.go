package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 2, 2, 1, 1}
	positive := firstMissingPositive(nums)
	fmt.Printf("res=%d", positive)
}

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	dest := 1
	prev := 0
	for _, num := range nums {
		if num <= 0 {
			continue
		}
		if num == prev {
			continue
		}
		prev = num
		if dest != num {
			return dest
		}
		dest++
	}
	return dest
}
