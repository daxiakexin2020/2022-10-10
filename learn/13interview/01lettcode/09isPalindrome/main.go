package main

import (
	"fmt"
	"strconv"
)

func main() {
	data := 0
	palindrome := isPalindrome(data)
	fmt.Printf("resL=%t", palindrome)
}

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i := 0; i <= len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}
