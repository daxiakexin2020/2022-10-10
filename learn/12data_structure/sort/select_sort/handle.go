package select_sort

// O(n2) 选择排序
func Handle(list []int) {
	/*
		选择排序，一般我们指的是简单选择排序，也可以叫直接选择排序，它不像冒泡排序一样相邻地交换元素，而是通过选择最小的元素，每轮迭代只需交换一次。虽然交换次数比冒泡少很多，但效率和冒泡排序一样的糟糕。

		选择排序属于选择类排序算法。

		我打扑克牌的时候，会习惯性地从左到右扫描，然后将最小的牌放在最左边，然后从第二张牌开始继续从左到右扫描第二小的牌，放在最小的牌右边，以此反复。选择排序和我玩扑克时的排序特别相似。

	*/

	n := len(list)
	for i := 0; i < n-1; i++ {
		min := list[i]
		minIndex := i
		for j := i + 1; j < n; j++ {
			if list[j] < min {
				min = list[j]
				minIndex = j
			}
		}
		if minIndex != i {
			list[i], list[minIndex] = list[minIndex], list[i]
		}
	}
}
