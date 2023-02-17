package main

import "01lettcode/helper"

func main() {
	l1 := helper.MakeList(3, helper.ASC)
	l2 := helper.MakeList(4, helper.ASC)

	helper.ShowList(l1)
	helper.ShowList(l2)

	res := addTwoNumbers(l1, l2) //321+4321 = 4642    2->4->6->4
	helper.ShowList(res)
}

func addTwoNumbers(l1 *helper.ListNode, l2 *helper.ListNode) *helper.ListNode {
	/*
		给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

		请你将两个数相加，并以相同形式返回一个表示和的链表。

		你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

		例如：
		输入：l1 = [2,4,3], l2 = [5,6,4]
		输出：[7,0,8]
		解释：342 + 465 = 807.

		来源：力扣（LeetCode）
		链接：https://leetcode.cn/problems/add-two-numbers
		著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
	*/
	var head *helper.ListNode
	var tail *helper.ListNode
	var carry int

	for l1 != nil || l2 != nil {
		var n1 int
		var n2 int

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}

		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &helper.ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &helper.ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &helper.ListNode{Val: carry}
	}

	return head
}
