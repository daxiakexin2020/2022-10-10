package main

import "fmt"

/*
*
给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
*/
func main() {
	nums := []int{1, 2, 3, 1, 4}
	res := findDuplicate(nums)
	fmt.Printf("res=%d", res)
}

func findDuplicate(nums []int) int {
	return 0
}
