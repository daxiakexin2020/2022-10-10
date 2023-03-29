package main

import (
	"01lettcode/helper"
	"fmt"
)

func main() {
	/**
	  	给定一个长度为 n 的链表 head

	    对于列表中的每个节点，查找下一个 更大节点 的值。也就是说，对于每个节点，找到它旁边的第一个节点的值，这个节点的值 严格大于 它的值。

	    返回一个整数数组 answer ，其中 answer[i] 是第 i 个节点( 从1开始 )的下一个更大的节点的值。如果第 i 个节点没有下一个更大的节点，设置 answer[i] = 0 。

	*/
	list := helper.MakeList(7, helper.ASC)
	helper.ShowList(list)
	res := nextLargerNodes(list)
	fmt.Println("res:", res)
}

func nextLargerNodes(head *helper.ListNode) []int {
	var values []int
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	var ret []int
	for i := 0; i < len(values); i++ {
		var flag bool
		for j := i + 1; j < len(values); j++ {
			if values[j] > values[i] {
				ret = append(ret, values[j])
				flag = true
				break
			}
		}
		if !flag {
			ret = append(ret, 0)
		}
	}
	return ret
}
