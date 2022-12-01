package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "abcdab"
	res := lengthOfLongestSubstring(s)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	/*
		给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

		例如：
		输入: s = "abcabcbb"
		输出: 3
		解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

		解题思路
		核心：只增大不减小的滑动窗口
		流程：两个指针start和end表示窗口大小，遍历一次字符串，窗口在遍历过程中滑动或增大
		tips：配合画图思考更佳

		窗口内没有重复字符：此时判断i+1与end的关系，超过表示遍历到窗口之外了，增大窗口大小
		窗口内出现重复字符：此时两个指针都增大index+1，滑动窗口位置到重复字符的后一位
		遍历结束，返回end-start，窗口大小
		思考：如果需要返回字符串怎么做？
		解答：只需要在窗口增大的时候记录start指针即可

	*/

	start, end := 0, 0
	for i := 0; i < len(s); i++ {

		//第一次循环
		//0：0 截取s的0：0，第一位，第一位

		//第二次
		//1：1    1

		index := strings.Index(s[start:i], string(s[i]))

		fmt.Println("i:", i, "start:", start, "end:", end, "s[start:i]:", s[start:i], "string(s[i]:", string(s[i]), "index:", index)

		//不存在
		if index == -1 {
			//0+1 >0  => end = 0+1   end后移动一位
			if i+1 > end {
				end = i + 1
			}
		} else {
			//存在     start = 0+1  start后移动一位   end后移动一位
			start += index + 1 //start=1  start=2
			end += index + 1   //end=1     end=2
		}
	}
	return end - start
}
