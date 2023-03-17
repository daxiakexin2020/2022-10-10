package lru

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type ListNode struct {
	head     *node
	tail     *node
	capacity int64
	len      int64
	mu       sync.RWMutex
}

type node struct {
	value      string
	createTime int64
	next       *node
}

const DefaultListNodeCapacity = 2 << 10

var defaultListNode = NewListNode(DefaultListNodeCapacity)

func DefaultListNode() *ListNode {
	return defaultListNode
}

func NewListNode(capacity int64) *ListNode {
	l := &ListNode{
		head:     newNode("0"),
		tail:     newNode("1"),
		capacity: capacity,
	}
	l.head.next = l.tail
	return l
}

func (ln *ListNode) AddNodeToTail(node *node) {
	ln.mu.Lock()
	ln.mu.Unlock()
	ln.tail.next = node
	ln.tail = node
	if ln.head == nil { //头部删除没了，为nil，重置指针指向
		ln.head = ln.tail
	}
	atomic.AddInt64(&ln.len, 1)
}

func (ln *ListNode) RemoveHeadNode() {
	ln.mu.Lock()
	ln.mu.Unlock()
	if ln.head != nil {
		ln.head = ln.head.next
		atomic.AddInt64(&ln.len, -1)
	}
}

func newNode(roomID string) *node {
	return &node{value: roomID, createTime: time.Now().Unix()}
}

func (ln *ListNode) Add(value string) {
	ln.AddNodeToTail(newNode(value))
}

func (ln *ListNode) GetHeadValue(remove func(value string, addtime int64) bool) string {
	head := ln.head
	if head == nil {
		return ""
	}
	value := head.value
	ctime := head.createTime
	if remove(value, ctime) {
		ln.RemoveHeadNode()
		return value
	}
	return ""
}

func (l *ListNode) Show() {
	n := l.head
	for n != nil {
		fmt.Println("node value:", n.value)
		n = n.next
	}
}
