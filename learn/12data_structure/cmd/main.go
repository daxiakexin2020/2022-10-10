package main

import (
	"12data_structure/sort/maopao_sort"
	"12data_structure/sort/select_sort"
	"12data_structure/stack/array_stack"
	"12data_structure/stack/link_stack"
	"fmt"
	"time"
)

func main() {
	//testArrayStack()
	//testgoArrayStack()
	//testLinkStack()
	//testMaoPaoSort()
	testSelectSort()
}

func testArrayStack() {
	link := array_stack.NewArrayStack()
	link.Push("a")
	link.Push("b")
	fmt.Println("array stack size", link.Size())
	fmt.Println("array stack top", link.Top())
	fmt.Println("array stack top", link.Pop())
	fmt.Println("出栈以后")
	fmt.Println("array stack size", link.Size())
	fmt.Println("array stack top", link.Top())
}

func testgoArrayStack() {
	link := array_stack.NewArrayStack()
	link.Push("a")
	link.Push("b")
	link.Push("c")
	fmt.Println("出栈之前")
	fmt.Println("array stack size", link.Size())
	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println("array stack top", link.Pop())
		}()
	}
	time.Sleep(time.Second * 2)
	fmt.Println("出栈以后")
	fmt.Println("array stack size", link.Size())
	fmt.Println("array stack top", link.Top())
}

func testLinkStack() {
	link := link_stack.NewLinkStack()
	link.Push("a")
	link.Push("b")
	link.Push("c")
	link.Show()

	fmt.Println("pop 元素", link.Pop())
	fmt.Println("pop 元素", link.Pop())
	fmt.Println("pop 元素", link.Pop())
}

func testMaoPaoSort() {
	list := []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}
	maopao_sort.Handle(list)
	fmt.Println("list", list)
}

func testSelectSort() {
	list := []int{1, 10, 2, 9, 3, 8, 4, 7, 5, 6}
	select_sort.Handle(list)
	fmt.Println("list", list)
}
