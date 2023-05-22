package main

import (
	"01lettcode/helper"
	"sort"
)

func main() {
	list := helper.MakeList(5, helper.DESC)
	sortList := insertionSortList(list)
	helper.ShowList(sortList)
}

func insertionSortList(head *helper.ListNode) *helper.ListNode {
	var vals []int
	for head != nil {
		vals = append(vals, head.Val)
		head = head.Next
	}

	retHead := &helper.ListNode{}
	curr := retHead
	sort.Ints(vals)

	for _, v := range vals {
		node := &helper.ListNode{
			Val:  v,
			Next: nil,
		}
		if curr == nil {
			curr.Val = v
		} else {
			curr.Next = node
			curr = curr.Next
		}
	}
	return retHead.Next
}
