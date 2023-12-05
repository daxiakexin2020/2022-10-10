package main

import "fmt"

func main() {

	/*
			知识：
				1.网络、计算机原理！
					每天晚上1-2个小时

				2.算法、数据结构、运维
					每天1道算法题

				3.k8s docker

		 		4.数据库   es、mongodb、redis

				5.python/java/c

				6.常见的架构内部原理：例如主主、主从、集群.....

			不能长时间刷头条
			每天坚持面1个单位
				每道面试题都要总结一下
			晚上坚持看书、看视频

		需要反复学习、动手、记忆
		如何成长？
			运维？
	*/

	test()
}

func test() {
	i := test2()
	fmt.Println("a", i)

	//test3()
	test4()
}

func test2() int {
	a := 11
	defer func() {
		a += 2
	}()
	a = 3
	return a
}

func test3() {
	ch := make(chan int)
	close(ch)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// 这里尝试向已关闭的channel发送数据，将会引发panic
	ch <- 1

	fmt.Println("This line will not be executed.")

	test2()
}

func deferAndReturn() (result int) {
	defer func() {
		// 这里将会改变返回值
		result += 10
	}()
	// 返回值被设置为0，但此时不是立即返回
	return 0
}

func test4() {
	value := deferAndReturn()
	fmt.Println(value) // 输出：10，而不是0，因为defer中修改了result
}
