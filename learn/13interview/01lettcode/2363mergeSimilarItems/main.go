package main

import (
	"fmt"
	"sort"
)

func main() {
	/**
	给你两个二维整数数组 items1 和 items2 ，表示两个物品集合。每个数组 items 有以下特质：

	items[i] = [valuei, weighti] 其中 valuei 表示第 i 件物品的 价值 ，weighti 表示第 i 件物品的 重量 。
	items 中每件物品的价值都是 唯一的 。
	请你返回一个二维数组 ret，其中 ret[i] = [valuei, weighti]， weighti 是所有价值为 valuei 物品的 重量之和 。

	注意：ret 应该按价值 升序 排序后返回。

	*/

	items2 := [][]int{
		[]int{1, 2},
		[]int{5, 6},
		[]int{3, 4},
		//[]int{2, 4},
		//[]int{6, 1},
	}

	items1 := [][]int{
		[]int{6, 7},
		[]int{7, 8},
		[]int{3, 1},
		[]int{8, 9},
	}

	res := mergeSimilarItems(items1, items2)
	fmt.Println(res)
}

// 双指针
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {

	sort.Slice(items1, func(i, j int) bool {
		return items1[i][0] < items1[j][0]
	})

	sort.Slice(items2, func(i, j int) bool {
		return items2[i][0] < items2[j][0]
	})

	var item1FlagIndex int
	var item2FlagIndex int
	var ret [][]int
L:
	for {
		item1Value := items1[item1FlagIndex][0]
		item1Weight := items1[item1FlagIndex][1]
		item2Value := items2[item2FlagIndex][0]
		item2Weight := items2[item2FlagIndex][1]
		var new []int
		if item1Value == item2Value {
			new = []int{item2Value, item1Weight + item2Weight}
			item1FlagIndex++
			item2FlagIndex++
		}
		if item1Value < item2Value {
			new = []int{item1Value, item1Weight}
			item1FlagIndex++
		}
		if item1Value > item2Value {
			new = []int{item2Value, item2Weight}
			item2FlagIndex++
		}
		ret = append(ret, new)

		if item1FlagIndex == len(items1) {
			ret = append(ret, items2[item2FlagIndex:]...)
			break L
		}
		if item2FlagIndex == len(items2) {
			ret = append(ret, items1[item1FlagIndex:]...)
			break L
		}
	}
	return ret
}

// map 相对耗内存
func mergeSimilarItemsMap(item1, item2 [][]int) [][]int {
	mp := map[int]int{}
	for _, a := range item1 {
		mp[a[0]] += a[1]
	}
	for _, a := range item2 {
		mp[a[0]] += a[1]
	}
	var ans [][]int
	for a, b := range mp {
		ans = append(ans, []int{a, b})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}
