package tree

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Data  string
	Left  *TreeNode
	Right *TreeNode
}

/**
遍历树，一般有4种方式
先序遍历：先访问根节点，再访问左子树，最后访问右子树。
中序遍历：先访问左子树，再访问根节点，最后访问右子树。
后序遍历：先访问左子树，再访问右子树，最后访问根节点。
层次遍历：每一层从左到右访问每一个节点。
*/

// 先序遍历 ：先访问根节点，再访问左子树，最后访问右子树。
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	fmt.Println(tree.Data, "")
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

// 中序遍历：先访问左子树，再访问根节点，最后访问右子树。
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	//todo 递归，后进先出，所以，应该是从树的底层向上遍历
	MidOrder(tree.Left)
	fmt.Println(tree.Data, "")
	MidOrder(tree.Right)
}

// 后序遍历：先访问左子树，再访问右子树，最后访问根节点。
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}
	MidOrder(tree.Left)
	MidOrder(tree.Right)
	fmt.Println(tree.Data, "")
}

/**
层次遍历：每一层从左到右访问每一个节点。
层次遍历较复杂，用到一种名叫广度遍历的方法，需要使用辅助的先进先出的队列。

1、先将树的根节点放入队列。
2、从队列里面 remove 出节点，先打印节点值，如果该节点有左子树节点，左子树入栈，如果有右子树节点，右子树入栈。
3、重复2，直到队列里面没有元素。
*/

func LayerOrder(treeNode *TreeNode) {
	if treeNode == nil {
		return
	}

	// 新建队列
	queue := new(LinkQueue)

	// 根节点先入队
	queue.Add(treeNode)
	for queue.size > 0 {
		// 不断出队列
		element := queue.Remove()

		// 先打印节点值
		fmt.Print(element.Data, " ")

		// 左子树非空，入队列
		if element.Left != nil {
			queue.Add(element.Left)
		}

		// 右子树非空，入队列
		if element.Right != nil {
			queue.Add(element.Right)
		}
	}
}

// 链表节点
type LinkNode struct {
	Next  *LinkNode
	Value *TreeNode
}

// 链表队列，先进先出
type LinkQueue struct {
	root *LinkNode  // 链表起点
	size int        // 队列的元素数量
	lock sync.Mutex // 为了并发安全使用的锁
}

// 入队
func (queue *LinkQueue) Add(v *TreeNode) {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 如果栈顶为空，那么增加节点
	if queue.root == nil {
		queue.root = new(LinkNode)
		queue.root.Value = v
	} else {
		// 否则新元素插入链表的末尾
		// 新节点
		newNode := new(LinkNode)
		newNode.Value = v

		// 一直遍历到链表尾部
		nowNode := queue.root
		for nowNode.Next != nil {
			nowNode = nowNode.Next
		}

		// 新节点放在链表尾部
		nowNode.Next = newNode
	}

	// 队中元素数量+1
	queue.size = queue.size + 1
}

// 出队
func (queue *LinkQueue) Remove() *TreeNode {
	queue.lock.Lock()
	defer queue.lock.Unlock()

	// 队中元素已空
	if queue.size == 0 {
		panic("over limit")
	}

	// 顶部元素要出队
	topNode := queue.root
	v := topNode.Value

	// 将顶部元素的后继链接链上
	queue.root = topNode.Next

	// 队中元素数量-1
	queue.size = queue.size - 1

	return v
}

// 队列中元素数量
func (queue *LinkQueue) Size() int {
	return queue.size
}
