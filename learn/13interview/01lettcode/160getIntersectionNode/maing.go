package main

import (
	"01lettcode/helper"
	"fmt"
)

func main() {
	listA := helper.MakeList(3)
	listB := helper.MakeList(5)
	res := getIntersectionNode(listA, listB)
	helper.ShowList(res)
	fmt.Println("res=%v", res)
}

func getIntersectionNode(headA, headB *helper.ListNode) *helper.ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA == nil { // 如果第一次遍历到链表尾部，就指向另一个链表的头部，继续遍历，这样会抵消长度差。如果没有相交，因为遍历长度相等，最后会是 nil ==  nil
			curA = headB
		} else {
			curA = curA.Next
		}
		if curB == nil {
			curB = headA
		} else {
			curB = curB.Next
		}
	}
	return curA

}
