package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, -1, -6, 10, 20}
	res := maxSubArray(nums)
	fmt.Printf("res=%d", res)
}

func maxSubArray(nums []int) int {
	/**
	给你一个整数数组 nums ，请你找出一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

	子数组 是数组中的一个连续部分。

	例如：
	输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出：6
	解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。
	*/
	var temSum int
	maxSum := nums[0]
	for i := 0; i < len(nums); i++ {
		temSum += nums[i]
		if temSum > maxSum {
			maxSum = temSum
		}
		if temSum < 0 {
			temSum = 0
		}
	}
	return maxSum
}
