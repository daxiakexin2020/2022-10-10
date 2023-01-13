package helper

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type sort_type string

const (
	DESC sort_type = "desc"
	ASC  sort_type = "asc"
)

func MakeList(len int, tsort sort_type) *ListNode {
	var tail *ListNode
	var head *ListNode

	m := func(v int) {
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

	if tsort == DESC {
		for v := len; v >= 1; v-- {
			m(v)
		}
	} else {
		for v := 1; v <= len; v++ {
			m(v)
		}
	}
	return head
}

func ShowList(node *ListNode) {
	fmt.Println("**********************showList start**********************")
	for node != nil {
		fmt.Printf("value=%d\n", node.Val)
		node = node.Next
	}
}
