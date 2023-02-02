package main

import "01lettcode/helper"

func main() {
	list := helper.MakeList(5, helper.DESC)
	helper.ShowList(list)
	res := removeNodes(list)
	helper.ShowList(res)
}

func removeNodes(head *helper.ListNode) *helper.ListNode {
	/**
	单调栈
	*/
	stack := []int{}
	for head != nil {
		//比较当前value与栈顶的元素的大小（切片最后一个元素），如果当前元素大于栈顶元素，则栈顶元素出栈
		for len(stack) > 0 && head.Val > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, head.Val)
		head = head.Next
	}
	head = &helper.ListNode{}
	node := head
	for _, v := range stack {
		head.Next = &helper.ListNode{Val: v}
		head = head.Next
	}
	return node.Next
}
