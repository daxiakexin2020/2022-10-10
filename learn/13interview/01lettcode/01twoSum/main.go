package main

import "fmt"

func main() {
	nums := []int{1, 10, 2, 5, 3, 4}
	target := 5
	res := twoSum(nums, target)
	fmt.Printf("res:=%v", res)
}

func twoSum(nums []int, target int) []int {
	/**
	给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
	你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
	你可以按任意顺序返回答案。
	例如：
	输入：nums = [3,2,4], target = 6
	输出：[1,2]

	来源：力扣（LeetCode）
	链接：https://leetcode.cn/problems/two-sum
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
	*/

	tmpMap := make(map[int]int)
	for index, value := range nums {
		if i, ok := tmpMap[value]; ok {
			return []int{i, index}
		} else {
			tmpMap[target-value] = index
		}
	}
	return nil
}