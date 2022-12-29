package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 3, 4, 4, 4, 4, 4}
	res := majorityElement(nums)
	fmt.Printf("res=%d", res)
}

/*
*

	给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
	你可以假设数组是非空的，并且给定的数组总是存在多数元素。
*/
func majorityElement(nums []int) int {
	tmp := make(map[int]int)
	var max int
	for _, num := range nums {
		tmp[num]++
		tmpVal := tmp[num]
		if tmpVal > len(nums)/2 && tmpVal > max {
			max = num
		}
	}
	return max
}
