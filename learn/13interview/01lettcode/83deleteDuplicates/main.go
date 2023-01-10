package main

import "01lettcode/helper"

func main() {
	list := helper.MakeList(5)
	helper.ShowList(list)
	res := deleteDuplicates(list)
	helper.ShowList(res)
}

func deleteDuplicates(head *helper.ListNode) *helper.ListNode {
	if head == nil {
		return head
	}
	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
