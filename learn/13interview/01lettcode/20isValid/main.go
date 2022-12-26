package main

import (
	"fmt"
)

func main() {
	str := "(){}[{}]"
	res := isValid(str)
	fmt.Printf("res=%v", res)
}

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}
	pairs := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	var stack []byte

	for i := 0; i < n; i++ {
		//开始的括号
		startByte := pairs[s[i]] //找到对应的开始结束字符关系
		if startByte <= 0 {
			//todo 可能是开始的3种字符，也可能是其他字符   压栈
			stack = append(stack, s[i]) // todo ( 、{ 、[  value不在map中，说明是开始的3种字符或者是其他字符....
		} else {
			//栈为空，或者是栈顶的字符与要比较的字符不相等，则说明抵消不了，不能出栈
			if len(stack) == 0 || stack[len(stack)-1] != startByte {
				return false
			}
			//TODO 出栈
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
