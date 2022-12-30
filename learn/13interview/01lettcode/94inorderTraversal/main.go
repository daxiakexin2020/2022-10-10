package main

import (
	"01lettcode/helper"
	"fmt"
)

func main() {
	root := helper.MakeTree(3)
	res := inorderTraversal(root)
	fmt.Println("res", res)
}

func inorderTraversal(root *helper.TreeNode) (res []int) {
	var inorder func(node *helper.TreeNode)
	inorder = func(node *helper.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}
