package main

import (
	"fmt"
	"sync"
)

func main() {
	s := Constructor()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	fmt.Println(s.Data, s.Pop, s.Data)
}

type SortedStack struct {
	Data []int
	mu   sync.Mutex
}

func Constructor() SortedStack {
	return SortedStack{
		Data: make([]int, 0),
		mu:   sync.Mutex{},
	}
}

func (this *SortedStack) Push(val int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	//如果栈为空，或者是栈顶元素小于目标元素，直接在栈顶追加元素
	l := len(this.Data)
	if l == 0 || this.Data[l-1] < val {
		this.Data = append(this.Data, val)
		return
	}
	//寻找应该插入的位置
	for l != 0 && this.Data[l-1] > val {
		l--
	}
	right := append([]int{val}, this.Data[l:]...)
	left := this.Data[:l]
	this.Data = append(left, right...)
}

func (this *SortedStack) Pop() {
	if this.IsEmpty() {
		return
	}
	this.Data = this.Data[1:]
}

// 返回索引位置
func (this *SortedStack) Peek() int {
	if this.IsEmpty() {
		return -1
	}
	return this.Data[0]
}

func (this *SortedStack) IsEmpty() bool {
	this.mu.Lock()
	defer this.mu.Unlock()
	return len(this.Data) == 0
}
