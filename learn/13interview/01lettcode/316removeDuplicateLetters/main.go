package main

func main() {

}

func removeDuplicateLetters(s string) string {
	/**
	给你一个字符串 s ，请你去除字符串中重复的字母，使得每个字母只出现一次。需保证 返回结果的字典序最小（要求不能打乱其他字符的相对位置）。

	示例 1：
		输入：s = "bcabc"
		输出："abc"


	方法一：贪心 + 单调栈
	思路与算法
	*/

	//26个小写字母
	left := [26]int{}
	for _, ch := range s {
		//将每个出现的字符的个数++，统计一下，每个字符出现的个数，方便后续出栈统计
		left[ch-'a']++
	}
	stack := []byte{}

	//是否在栈中
	inStack := [26]bool{}
	for i := range s {
		ch := s[i]

		//次字符不在栈中
		if !inStack[ch-'a'] {

			//栈不为空，并且当前字符的字典小于栈顶的字符的字典
			for len(stack) > 0 && ch < stack[len(stack)-1] {
				last := stack[len(stack)-1] - 'a'
				//此元素已经是最后一个，后续不会再出现，直接退出循环
				if left[last] == 0 {
					break
				}
				//如果，后续还会出现此元素，并且当前元素小于栈顶元素，栈顶元素出栈
				stack = stack[:len(stack)-1]

				//此字符标记为，不在栈中
				inStack[last] = false
			}
			stack = append(stack, ch)
			inStack[ch-'a'] = true
		}
		left[ch-'a']--
	}
	return string(stack)
}
