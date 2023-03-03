package main

import (
	"fmt"
	"sort"
)

func main() {

	//data := []int{10, 1, 25, 11, 24, 17, 26, 2, 3, 4, 5} //sort
	//data := []int{1, 2, 3, 4, 5, 10, 11, 17, 24, 25, 26} //sort
	data := []int{10, 9, 8, 7, 6, 5, 2}

	//Search函数采用二分法搜索找到[0, n)区间内最小的满足f(i)==true的值i。也就是说，Search函数希望f在输入位于区间[0, n)的前面某部分（可以为空）时返回假，而在输入位于剩余至结尾的部分（可以为空）时返回真；Search函数会返回满足f(i)==true的最小值i。如果没有该值，函数会返回n。注意，未找到时的返回值不是-1，这一点和strings.Index等函数不同。Search函数只会用区间[0, n)内的值调用f。
	//一般使用Search找到值x在插入一个有序的、可索引的数据结构时，应插入的位置。这种情况下，参数f（通常是闭包）会捕捉应搜索的值和被查询的数据集。
	//例如，给定一个递增顺序的切片，调用Search(len(data), func(i int) bool { return data[i] >= 23 })会返回data中最小的索引i满足data[i] >= 23。如果调用者想要知道23是否在切片里，它必须另外检查data[i] == 23。
	search := sort.Search(len(data), func(i int) bool {
		//递增的slice
		//return data[i] >= 23

		//递减的slice,注意大于小于符号
		return data[i] > 7
	})
	fmt.Println("search:", search) //

	//将data排序为递增顺序
	sort.Ints(data)
	fmt.Println("Ints():", data)

	//检查data是否是递增顺序的
	sorted := sort.IntsAreSorted(data)
	fmt.Printf("IntsAreSorted=%t\n", sorted)

	//SearchInts在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。
	ints := sort.SearchInts(data, 7)
	fmt.Printf("SearchInts=%d\n", ints)

	//Strings函数将a排序为递增顺序。
	str := []string{"abcdefg ", "ihew", "a", "b", "e", "d", "c"}
	sort.Strings(str)
	fmt.Printf("str=%v\n", str) //str=[a abcdefg  b c d e ihew]

	//SearchStrings在递增顺序的a中搜索x，返回x的索引。如果查找不到，返回值是x应该插入a的位置（以保证a的递增顺序），返回值可以是len(a)。
	strings := sort.SearchStrings(str, "a")
	fmt.Printf("SearchStrings=%v\n", strings)
}
