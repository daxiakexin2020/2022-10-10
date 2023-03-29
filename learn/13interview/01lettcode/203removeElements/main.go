package main

import "01lettcode/helper"

func main() {
	/**
	给你一个链表的头节点 head 和一个整数 val ，请你删除链表中所有满足 Node.val == val 的节点，并返回 新的头节点 。
	*/
	list := helper.MakeList(6, helper.ASC)
	helper.ShowList(list)
	elements := removeElements(list, 1)
	helper.ShowList(elements)
}

func removeElements(head *helper.ListNode, val int) *helper.ListNode {
	dummyHead := &helper.ListNode{Next: head}
	for tmp := dummyHead; tmp.Next != nil; {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummyHead.Next
}
