package node

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeList(len int) *ListNode {
	var tail *ListNode
	var head *ListNode
	for v := 1; v <= len; v++ {
		l := &ListNode{
			Val:  v,
			Next: nil,
		}
		if head == nil {
			head = l
			tail = head
		} else {
			tail.Next = l
			tail = tail.Next
		}
	}
	return head
}

func ShowList(node *ListNode) {
	for node != nil {
		fmt.Printf("value=%d\n", node.Val)
		node = node.Next
	}
}
