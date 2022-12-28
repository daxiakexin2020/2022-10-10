package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	res := containsDuplicate(nums)
	fmt.Printf("res=%t", res)
}

func containsDuplicate(nums []int) bool {
	tmp := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := tmp[num]; ok {
			return true
		}
		tmp[num] = struct{}{}
	}
	return false
}
