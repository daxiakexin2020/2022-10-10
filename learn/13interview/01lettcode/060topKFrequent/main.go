package main

import (
	"fmt"
	"sort"
)

func main() {
	/**
	给定一个整数数组 nums 和一个整数 k ，请返回其中出现频率前 k 高的元素。可以按 任意顺序 返回答案。
	*/
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	res := topKFrequent(nums, 10)
	fmt.Println("res:", res) //[1 2 5 3 6 7 8 4 10 11]   [1,2,5,3,6,7,4,8,10,11]
}

func topKFrequent(nums []int, k int) []int {
	var counts []int
	numCount := make(map[int]int)
	countSet := make(map[int][]int)

	for _, num := range nums {
		numCount[num] += 1
	}
	for num, count := range numCount {
		countSet[count] = append(countSet[count], num)
		counts = append(counts, count)
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	var ret []int
	lastCount := -1
	for _, count := range counts {
		if count == lastCount { //去重
			continue
		}
		if len(ret) == k {
			return ret
		}
		numSet := countSet[count]
		if len(numSet)+len(ret) <= k {
			ret = append(ret, numSet...)
		} else {
			needLen := k - len(ret)
			needs := numSet[0:needLen]
			ret = append(ret, needs...)
		}
		lastCount = count
	}
	return ret
}
