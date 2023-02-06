package main

import "fmt"

/**
递归
	递归的执行流程和栈一样的，都是后进先出  类似手枪弹夹
	运动开始时，首先为递归调用建立一个工作栈，其结构包括值参、局部变量和返回地址；
	每次执行递归调用之前，把递归函数的值参、局部变量的当前值以及调用后的返回地址压栈；
	每次递归调用结束后，将栈顶元素出栈，使相应的值参和局部变量恢复为调用前的值，然后转向返回地址指定的位置继续执行。
*/

func main() {
	//1 1 2 3 5 8 13 21
	//res := recursion(5)
	//fmt.Printf("**************res=%d*****************\n", res)

	//1,2,3,4,5
	//res := recursion01(5)
	//fmt.Printf("**************res=%d*****************\n", res)

	//1,2,3,4,5
	res := recursion02(5)
	fmt.Printf("**************res=%d*****************\n", res)
}

func recursion(n int) int {
	if n < 2 {
		return n
	}
	//n=4, 只走f1，直接去递归，直到n<2 退出  第一次被调用3次，4，3，2
	//fmt.Printf("f1前：n=%d, n-1=%d\n", n, n-1)
	f1 := recursion(n - 1) //f1 接收返回值

	//TODO 后进先出   f2的n，应该是2，3，4的顺序
	//fmt.Printf("f1后，f2前：n=%d, n-1=%d, f1=%d\n", n, n-1, f1)
	f2 := recursion(n - 2) //f2 接收返回值

	res := f1 + f2
	fmt.Printf("f1后，f2后,n=%d,f1=%d,f2=%d,res=%d\n", n, f1, f2, res)

	//todo res又作为返回值返回给了调用者，作为f1或者是f2，累加的结果，就是从这里出去的
	return res
}

func recursion01(n int) int {
	if n <= 0 {
		return 1
	}
	return n * recursion01(n-1)
}

func recursion02(n int) int {
	if n <= 0 {
		return 0
	}
	return n + recursion02(n-1)
}
