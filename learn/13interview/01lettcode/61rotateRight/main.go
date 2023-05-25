package main

import (
	"01lettcode/helper"
	"fmt"
)

func main() {
	list := helper.MakeList(5, helper.ASC)
	right := rotateRight(list, 2)
	helper.ShowList(right)
}

func rotateRight(head *helper.ListNode, k int) *helper.ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	item := head
	for item.Next != nil {
		item = item.Next
		n++
	}

	add := n - k%n //2%5
	fmt.Println("info:", n, k%n, add)
	//旋转了一圈，不用动
	if add == n {
		return head
	}
	item.Next = head //形成闭环
	for add > 0 {
		item = item.Next
		add--
	}
	ret := item.Next
	item.Next = nil
	return ret
}
