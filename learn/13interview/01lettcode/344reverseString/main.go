package main

import "fmt"

func main() {
	/**
	编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。

	不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。

	*/
	s := []byte{'a', 'b', 'c', 'd'}
	reverseString(s)
	fmt.Println(string(s))
}

func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
}
