package main

import (
	"01lettcode/helper"
	"log"
)

func main() {
	list := helper.MakeList(5, helper.ASC)
	palindrome := isPalindrome(list)
	log.Printf("res:=%t\n", palindrome)
}

func isPalindrome(head *helper.ListNode) bool {

	if head == nil {
		return false
	}

	var stack []int
	cur := head
	for cur != nil {
		stack = append(stack, cur.Val)
		cur = cur.Next
	}
	for i := len(stack) - 1; i >= 0; i-- {
		val := stack[i]
		if head.Next == nil {
			return false
		}
		if val != head.Val {
			return false
		}
		head = head.Next
	}
	return true
}
