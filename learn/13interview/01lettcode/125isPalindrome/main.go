package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	s := "abc9A,A9cba"
	palindrome := isPalindrome(s)
	fmt.Println("Res:", palindrome)
}

func isPalindrome(s string) bool {
	/**
	如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
	字母和数字都属于字母数字字符。
	给你一个字符串 s，如果它是 回文串 ，返回 true ；否则，返回 false 。
	*/
	parrent := "[^0-9a-zA-Z]"
	r := regexp.MustCompile(parrent)
	findString := r.ReplaceAllString(s, "")
	lowerString := strings.ToLower(findString)
	var check []byte
	for i := 0; i < len(lowerString); i++ {
		check = append(check, lowerString[i])
	}
	for i := 0; i < len(lowerString)/2; i++ {
		if lowerString[i] != check[len(check)-i-1] {
			return false
		}
	}
	return true
}
