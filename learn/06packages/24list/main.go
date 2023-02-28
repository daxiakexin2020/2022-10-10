package main

import (
	"container/list"
	"fmt"
)

func main() {

	//创建一个新的双向链表
	l := list.New()
	fmt.Println("list.New:", l)

	//PushFront将一个值为v的新元素插入链表的第一个位置，返回生成的新元素。
	front := l.PushFront(1)
	fmt.Println("pushFront:", front)

	front2 := l.PushFront(2)
	fmt.Println("pushFront2:", front2)

	front3 := l.PushFront(3)
	fmt.Println("pushFront3:", front3)

	fmt.Println("l:", l)

	//返回链表的第一个元素或者nil
	k := l.Front()
	for k != nil {
		fmt.Println("value:", k.Value)
		k = k.Next()
	}

}
