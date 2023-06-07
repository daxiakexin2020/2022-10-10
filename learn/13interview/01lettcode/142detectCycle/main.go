package main

import "01lettcode/helper"

func main() {
	list := helper.MakeList(5, helper.ASC)
	cycle := detectCycle(list)
	helper.ShowList(cycle)
}

func detectCycle(head *helper.ListNode) *helper.ListNode {
	flags := map[*helper.ListNode]struct{}{}

	for head != nil {
		if _, ok := flags[head]; !ok {
			flags[head] = struct{}{}
			head = head.Next
		} else {
			return head
		}
	}
	return nil
}
