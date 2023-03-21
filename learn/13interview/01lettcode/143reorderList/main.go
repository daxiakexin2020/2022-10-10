package main

import (
	"01lettcode/helper"
)

func main() {
	list := helper.MakeList(5, helper.ASC)
	helper.ShowList(list)
	reorderList(list)
	helper.ShowList(list)
}

func reorderList(head *helper.ListNode) {
	/**
	给定一个单链表 L 的头节点 head ，单链表 L 表示为：
	L0 → L1 → … → Ln - 1 → Ln
	请将其重新排列后变为：
	L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …

	解法：
	因为链表不支持下标访问，所以我们无法随机访问链表中任意位置的元素。
	因此比较容易想到的一个方法是，我们利用线性表存储该链表，然后利用线性表可以下标访问的特点，直接按顺序访问指定元素，重建该链表即可。
	*/
	if head == nil {
		return
	}
	nodes := []*helper.ListNode{}
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}
	var i int
	j := len(nodes) - 1

	//nodes=>  1,2,3,4,5; 2,3,4,5; 3,4,5;  4,5; 5;
	// 0, 5
	for i < j {
		nodes[i].Next = nodes[j]
		//0=>1
		i++
		if i == j {
			break
		}
		// 5    = 1
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil
}
