package main

import "01lettcode/helper"

func main() {
	l1 := helper.MakeList(3, helper.ASC)
	l2 := helper.MakeList(3, helper.ASC)
	res := getIntersectionNode(l1, l2)
	helper.ShowList(res)
}

func getIntersectionNode(headA, headB *helper.ListNode) *helper.ListNode {
	flags := make(map[*helper.ListNode]struct{})
	for headA != nil {
		flags[headA] = struct{}{}
		headA = headA.Next
	}
	for headB != nil {
		if _, ok := flags[headB]; ok {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
