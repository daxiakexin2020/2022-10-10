package main

import (
	"math"
	"sync"
)

func main() {
	/**
	设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

	实现 MinStack 类:

	MinStack() 初始化堆栈对象。
	void push(int val) 将元素val推入堆栈。
	void pop() 删除堆栈顶部的元素。
	int top() 获取堆栈顶部的元素。
	int getMin() 获取堆栈中的最小元素。

	*/
}

type MinStack struct {
	stack    []int
	minStack []int
	mu       sync.Mutex
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minStack: []int{math.MaxInt64},
		mu:       sync.Mutex{},
	}
}

func (this *MinStack) Push(val int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.stack = append(this.stack, val)
	top := this.minStack[len(this.minStack)-1]
	this.minStack = append(this.minStack, min(val, top))
}

func (this *MinStack) Pop() {
	this.mu.Lock()
	defer this.mu.Unlock()
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
