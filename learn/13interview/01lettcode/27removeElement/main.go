package main

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	val := 0
	removeElement(nums, val)
}

func removeElement(nums []int, val int) int {
	/*
		给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

		不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

		元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

		例如：
		输入：nums = [3,2,2,3], val = 3
		输出：2, nums = [2,2]
		解释：函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。你不需要考虑数组中超出新长度后面的元素。例如，函数返回的新长度为 2 ，而 nums = [2,2,3,3] 或 nums = [2,2,0,0]，也会被视作正确答案。

	*/

	for i := 0; i < len(nums); {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
			continue
		}
		i++
	}
	return len(nums)
}
