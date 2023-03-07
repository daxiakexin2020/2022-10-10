package main

import (
	"fmt"
	"sort"
)

func main() {
	/**
	给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
	*/
	nums := []int{0, 1}
	res := missingNumber(nums)
	fmt.Printf("res:=%d\n", res)
}

func missingNumber(nums []int) int {
	sort.Ints(nums)
	dest := nums[0]
	if dest != 0 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num-dest == 1 {
			dest = num
		} else {
			return num - 1
		}
	}
	return dest + 1
}
