package maopao_sort

// O(n2)   O(n)    冒泡排序
func Handle(list []int) {
	n := len(list)
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if list[j] > list[j+1] {
				/*
					很多编程语言不允许这样：list[j], list[j+1] = list[j+1], list[j]，会要求交换两个值时必须建一个临时变量 a 来作为一个过渡，如：
					a := list[j+1]
					list[j+1] = list[j]
					list[j] = a
					但是 Golang 允许我们不那么做，它会默认构建一个临时变量来中转。
				*/
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
