package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 10}
	target := 2
	res := searchInsert(nums, target)
	fmt.Printf("res=%d", res)
}

func searchInsert(nums []int, target int) int {

	/*
		给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

		请必须使用时间复杂度为 O(log n) 的算法。

		例如：
		输入: nums = [1,3,5,6], target = 2
		输出: 1

	*/
	for index, num := range nums {
		if num >= target {
			return index
		}
	}
	return len(nums)
}
