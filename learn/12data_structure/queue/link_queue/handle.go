package link_queue

import (
	"fmt"
	"sync"
)

type LinkStack struct {
	root *LinkNode
	size int
	lock sync.Mutex
}

type LinkNode struct {
	Next  *LinkNode
	Value string
}

func NewLinkStack() *LinkStack {
	return &LinkStack{
		lock: sync.Mutex{},
	}
}

func (ls *LinkStack) Push(v string) {
	ls.lock.Lock()
	defer ls.lock.Unlock()

	if ls.root == nil {
		ls.root = &LinkNode{Value: v}
	} else {
		newNode := &LinkNode{Value: v}
		oldLink := ls.root
		for oldLink != nil {
			oldLink = oldLink.Next
		}
		oldLink.Next = newNode //插入到尾部
	}
	ls.size = ls.size + 1
}

func (ls *LinkStack) Pop() string {

	ls.lock.Lock()
	defer ls.lock.Unlock()
	if ls.IsEmpty() {
		panic("empty")
	}

	topNode := ls.root
	v := topNode.Value
	ls.root = topNode.Next //断开，把top断开
	ls.size = ls.size - 1
	return v
}

func (ls *LinkStack) Top() string {
	ls.lock.Lock()
	defer ls.lock.Unlock()
	if ls.IsEmpty() {
		panic("empty")
	}
	return ls.root.Value
}

func (ls *LinkStack) Size() int {
	return ls.size
}

func (ls *LinkStack) IsEmpty() bool {
	return ls.size == 0
}

func (ls *LinkStack) Show() {
	link := ls.root
	for link != nil {
		fmt.Println("**************元素**************", link.Value)
		link = link.Next
	}
}
