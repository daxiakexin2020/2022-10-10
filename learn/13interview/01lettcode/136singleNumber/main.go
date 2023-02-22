package main

import (
	"fmt"
)

/*
*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。

你必须设计并实现线性时间复杂度的算法来解决此问题，且该算法只使用常量额外空间。
*/
func main() {
	nums := []int{-1, 4, 4, 1, 2, 1, 2}
	res := singleNumber(nums)
	fmt.Printf("res=%d", res)
}

/**
任何数和0做异或运算，结果仍然是原来的数，即 a⊕0=a
任何数和其自身做异或运算，结果是0 即a⊕a=0
异或运算满足交换律和结合律,即 a⊕b⊕a=b⊕a⊕a=b⊕(a⊕a)=b⊕0=b
*/

func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
