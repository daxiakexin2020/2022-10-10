package service

import "fmt"

/**
坑位一：切片 dst 需要先初始化长度
不是你定义好类型，就能将 src 完全 copy 到 dst 的，你需要初始化长度。

如果 dst 长度小于 src 的长度，则 copy 部分；
如果大于，则全部拷贝过来，只是没占满 dst 的坑位而已；
相等时刚好不多不少 copy 过来。

坑位二: 源切片中元素类型为引用类型时，拷贝的是引用
由于只 copy 切片中的元素，所以如果切片元素的类型是引用类型，那么 copy 的也将是个引用。

*/

func doCopy() {
	b := make([]int, 4)
	a := []int{1, 2, 3}
	n := copy(b, a)
	a = append(a, 4)
	fmt.Println(b, n, a)
}
