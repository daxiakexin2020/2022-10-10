package double_link

import "sync"

type DoubleList struct {
	head *ListNode
	tail *ListNode
	len  int
	lock sync.Mutex
}

type ListNode struct {
	prev  *ListNode
	next  *ListNode
	value int
}

func (node *ListNode) GetValue() int {
	return node.value
}

func (node *ListNode) GetPrev() *ListNode {
	return node.prev
}

func (node *ListNode) GetNext() *ListNode {
	return node.next
}

func (node *ListNode) HasNext() bool {
	return node.next != nil
}

func (node *ListNode) HasPrev() bool {
	return node.prev != nil
}

func (node *ListNode) IsNil() bool {
	return node == nil
}
