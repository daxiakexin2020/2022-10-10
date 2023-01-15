package main

import (
	"01lettcode/helper"
)

func main() {
	head := helper.MakeList(3, helper.ASC)
	reverseListRes := reverseList(head)
	helper.ShowList(head)
	helper.ShowList(reverseListRes)
}

func reverseList(head *helper.ListNode) *helper.ListNode {
	/**
	给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

	*/
	//nil
	var prev *helper.ListNode

	//1,2,3
	curr := head

	for curr != nil {
		//curr:1   curr:2
		next := curr.Next // 2  3
		curr.Next = prev  //2=>nil
		prev = curr       //nil=>curr
		curr = next       //curr=>2
	}
	return prev

}
