package main

import (
	"fmt"
	"sort"
)

func main() {
	/**
	给你一个下标从 0 开始的整数数组 nums 。在一步操作中，你可以执行以下步骤：

	从 nums 选出 两个 相等的 整数
	从 nums 中移除这两个整数，形成一个 数对
	请你在 nums 上多次执行此操作直到无法继续执行。
	返回一个下标从 0 开始、长度为 2 的整数数组 answer 作为答案，其中 answer[0] 是形成的数对数目，answer[1] 是对 nums 尽可能执行上述操作后剩下的整数数目。
	*/
	nums := []int{1, 1}
	pairs := numberOfPairs(nums)
	fmt.Println(pairs)
}

func numberOfPairs(nums []int) []int {
	var doubleCount int
	leftoverCount := len(nums)
	if len(nums) <= 1 {
		return []int{doubleCount, len(nums)}
	}
	sort.Ints(nums)
	start := nums[0] - 1
	dest := start
	for i := 0; i < len(nums); i++ {
		if nums[i] == dest {
			doubleCount++
			leftoverCount = len(nums) - doubleCount*2
			dest = start
		} else {
			dest = nums[i]
		}
	}
	return []int{doubleCount, leftoverCount}
}
