package main

import (
	"fmt"
)

func main() {
	/**
	设计一个算法，算出 n 阶乘有多少个尾随零。

	示例 1:

	输入: 3
	输出: 0
	解释: 3! = 6, 尾数中没有零。
	*/

	n := 5
	res := trailingZeroes(n)
	fmt.Printf("res=%d", res)
}

func trailingZeroes(n int) (ans int) {
	for n > 0 {
		n /= 5
		ans += n
	}
	return
}
