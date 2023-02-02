package main

import "sync"

type CustomStack struct {
	stack []int
	mu    sync.Mutex
}

/*
*
请你设计一个支持对其元素进行增量操作的栈。

实现自定义栈类 CustomStack ：

CustomStack(int maxSize)：用 maxSize 初始化对象，maxSize 是栈中最多能容纳的元素数量。
void push(int x)：如果栈还未增长到 maxSize ，就将 x 添加到栈顶。
int pop()：弹出栈顶元素，并返回栈顶的值，或栈为空时返回 -1 。
void inc(int k, int val)：栈底的 k 个元素的值都增加 val 。如果栈中元素总数小于 k ，则栈中的所有元素都增加 val
*/
func Constructor(maxSize int) CustomStack {
	if maxSize < 0 {
		return CustomStack{}
	}
	return CustomStack{
		stack: make([]int, 0, maxSize),
		mu:    sync.Mutex{},
	}
}

func (this *CustomStack) Push(x int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.stack) == cap(this.stack) {
		return
	}
	this.stack = append(this.stack, x)
}

func (this *CustomStack) Pop() int {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.stack) == 0 {
		return -1
	}
	pop := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return pop
}

func (this *CustomStack) Increment(k int, val int) {
	this.mu.Lock()
	defer this.mu.Unlock()
	for i := 0; i < k && i < len(this.stack); i++ {
		if i >= k {
			break
		}
		this.stack[i] += val
	}
}
