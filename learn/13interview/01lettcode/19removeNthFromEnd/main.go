package main

import (
	"01lettcode/helper"
)

func main() {
	head := helper.MakeList(10)
	helper.ShowList(head)
	resList := removeNthFromEnd(head, 2)
	helper.ShowList(resList)
}

/*
*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
*/
func removeNthFromEnd(head *helper.ListNode, n int) *helper.ListNode {
	dummyHead := &helper.ListNode{} //空的
	dummyHead.Next = head           //将head挂在新的list上
	cur := head                     //参与遍历，当前的list
	prev := dummyHead               //保存上一个list
	i := 1
	for cur != nil {
		cur = cur.Next
		if i > n {
			prev = prev.Next //上一个节点=上一个下一个节点
		}
		i++
	}
	prev.Next = prev.Next.Next
	return dummyHead.Next
}
