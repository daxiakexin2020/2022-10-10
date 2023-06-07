package main

import "fmt"

func main() {
	//nums := []int{0, 1, 0, 3, 12}
	//nums := []int{0, 0, 1}
	nums := []int{1, 2, 3, 0, 0, 4, 5, 6, 0}
	moveZeroes(nums)
	fmt.Println("nums:", nums)
}

func moveZeroes(nums []int) {

	if len(nums) <= 1 {
		return
	}
	l := len(nums)
	for i := 0; i < l; i++ {
		num := nums[i]
		tmpIndex := i
		if num == 0 {
			for j := i; j < l-1; j++ {
				if nums[j+1] == 0 {
					continue
				}
				nums[tmpIndex] = nums[j+1]
				nums[j+1] = 0
				tmpIndex++
			}
			nums[l-1] = 0
		}
	}
}

func moveZeroes2(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
