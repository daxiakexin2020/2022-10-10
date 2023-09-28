package main

import (
	"01lettcode/helper"
)

func main() {
	list1 := helper.MakeList(3, helper.ASC)
	list2 := helper.MakeList(5, helper.ASC)
	res := mergeTwoLists(list1, list2)
	helper.ShowList(res)
}

func mergeTwoLists(l1 *helper.ListNode, l2 *helper.ListNode) *helper.ListNode {
	head := &helper.ListNode{}
	cur := head
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				cur.Next = l1
				l1 = l1.Next
			} else {
				cur.Next = l2
				l2 = l2.Next
			}
			cur = cur.Next
		} else if l1 != nil {
			cur.Next = l1
			break
		} else {
			cur.Next = l2
			break
		}
	}
	return head.Next
}
