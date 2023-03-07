package main

import (
	"01lettcode/helper"
	"fmt"
	"math/rand"
)

func main() {
	/**
	给你一个单链表，随机选择链表的一个节点，并返回相应的节点值。每个节点 被选中的概率一样 。
	实现 Solution 类：
	Solution(ListNode head) 使用整数数组初始化对象。
	int getRandom() 从链表中随机选择一个节点并返回该节点的值。链表中所有节点被选中的概率相等。
	*/
	head := helper.MakeList(5, helper.ASC)
	s := Constructor(head)
	fmt.Println("s:", s.GetRandom())
	fmt.Println("s:", s.GetRandom())
	fmt.Println("s:", s.GetRandom())
	fmt.Println("s:", s.GetRandom())
	fmt.Println("s:", s.GetRandom())
}

type Solution struct {
	indexs int
	list   map[int]*helper.ListNode
}

func Constructor(head *helper.ListNode) Solution {
	s := Solution{
		list: map[int]*helper.ListNode{},
	}
	var sort int
	for head != nil {
		s.list[sort] = head
		head = head.Next
		if head != nil {
			sort++
		}
	}
	s.indexs = sort
	return s
}

func (this *Solution) GetRandom() int {
	if this.indexs == 0 {
		return 0
	}
	//source := rand.NewSource(time.Now().UnixNano())
	//r := rand.New(source)
	intn := rand.Intn(this.indexs)
	if head, ok := this.list[intn]; ok {
		return head.Val
	}
	return 0
}
