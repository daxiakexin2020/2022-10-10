package helper

import (
	"math/rand"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewNode(n int) *TreeNode {
	return &TreeNode{Val: n}
}

func (nd *TreeNode) Search(dt int) *TreeNode {
	if nd == nil {
		return nil
	}
	//1
	if dt == nd.Val {
		return nd
	}

	//2 大于当前节点，递归右边
	if dt > nd.Val {
		return nd.Right.Search(dt)
	}
	//2 大于当前节点，递归左边
	if dt < nd.Val {
		return nd.Left.Search(dt)
	}
	return nil
}

func (nd *TreeNode) Insert(newNode *TreeNode) {
	//1
	if newNode.Val == nd.Val {
		return
	}
	//2
	if newNode.Val > nd.Val {
		if nd.Right == nil {
			nd.Right = newNode
		} else {
			nd.Right.Insert(newNode)
		}
	} else { //3 小于 继续比较插入到 左孩子
		if nd.Left == nil {
			nd.Left = newNode
		} else {
			nd.Left.Insert(newNode)
		}
	}
}

func MakeTree(n int) *TreeNode {
	root := NewNode(n)
	for i := 0; i < 2; i++ {
		n := rand.Intn(500)
		//fmt.Printf("i=%d,随机数=%d\n", i, n)
		root.Insert(NewNode(n))
	}
	return root
}
