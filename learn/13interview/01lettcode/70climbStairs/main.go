package main

import "fmt"

func main() {
	n := 4
	res := climbStairs(n)
	fmt.Printf("一共%d种方法", res)
}

/*
*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/
func climbStairs(n int) int {

	//动态规划
	var tempMap = make(map[int]int)
	tempMap[1] = 1
	tempMap[2] = 2
	for i := 3; i <= n; i++ {
		tempMap[i] = tempMap[i-1] + tempMap[i-2] //n=5  i=3 ,3=>2+1     ;  n=5  i=4, 4=>3+2=>5
	}
	return tempMap[n]
}
