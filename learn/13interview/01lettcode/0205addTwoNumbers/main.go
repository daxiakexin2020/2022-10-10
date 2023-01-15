package main

import (
	"01lettcode/helper"
	"fmt"
)

func main() {

	//654321
	//654321
	//1308642
	l1 := helper.MakeList(6, helper.ASC)
	l2 := helper.MakeList(6, helper.ASC)
	helper.ShowList(l1)
	helper.ShowList(l2)
	dest := addTwoNumbers(l1, l2)
	helper.ShowList(dest)
}

func addTwoNumbers(l1 *helper.ListNode, l2 *helper.ListNode) (head *helper.ListNode) {

	var remainder int
	var tail *helper.ListNode
	for l1 != nil || l2 != nil {
		l1v := 0
		l2v := 0
		if l1 != nil {
			l1v = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			l2v = l2.Val
			l2 = l2.Next
		}

		newv := (l1v + l2v + remainder) % 10
		fmt.Printf("l1v=%d,l2v=%d,remainder=%d,newv=%d\n", l1v, l2v, remainder, newv)
		remainder = int(l1v+l2v+remainder) / 10
		newNode := &helper.ListNode{Val: newv}

		if head == nil {
			head = newNode
			tail = head
		} else {
			tail.Next = newNode
			tail = tail.Next
		}
	}
	if remainder != 0 {
		tail.Next = &helper.ListNode{Val: remainder}
	}

	return
}
