package tree

import (
	"fmt"
	"testing"
)

func TestPreOrder(t *testing.T) {
	tree := MakeRree()
	fmt.Println("先序排序：")
	PreOrder(tree)
}

func TestMidOrder(t *testing.T) {
	tree := MakeRree()
	fmt.Println("中序遍历：")
	MidOrder(tree)
}

func TestPostOrder(t *testing.T) {
	tree := MakeRree()
	fmt.Println("后序遍历：")
	PostOrder(tree)
}

func TestLayerOrder(t *testing.T) {
	tree := MakeRree()
	fmt.Println("层次排序：")
	LayerOrder(tree)
}

func MakeRree() *TreeNode {
	tree := &TreeNode{Data: "A"}
	tree.Left = &TreeNode{Data: "B"}
	tree.Right = &TreeNode{Data: "C"}
	tree.Left.Left = &TreeNode{Data: "D"}
	tree.Left.Right = &TreeNode{Data: "E"}
	tree.Right.Left = &TreeNode{Data: "F"}
	return tree
}
