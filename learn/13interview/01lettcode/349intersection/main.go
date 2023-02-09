package main

import "fmt"

func main() {
	num1 := []int{1, 2, 3}
	num2 := []int{3, 3, 4, 5}
	res := intersection(num1, num2)
	fmt.Println("res", res)
}

func intersection(nums1 []int, nums2 []int) []int {
	/**
	给定两个数组 nums1 和 nums2 ，返回 它们的交集 。输出结果中的每个元素一定是 唯一 的。我们可以 不考虑输出结果的顺序 。
	*/
	result := make([]int, 0)
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}
	mnums1 := make(map[int]struct{}, len(nums1))
	flags := make(map[int]struct{}, len(nums1))
	for _, v := range nums1 {
		mnums1[v] = struct{}{}
	}
	for _, v := range nums2 {
		if _, ok := mnums1[v]; ok {
			if _, ok2 := flags[v]; !ok2 {
				result = append(result, v)
				flags[v] = struct{}{} //去重
			}
		}
	}
	return result
}
