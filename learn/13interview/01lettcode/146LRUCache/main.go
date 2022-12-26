package main

func main() {
	/**
	实现 LRUCache 类：
		LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
		int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
		void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。
	    如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
		函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

	哈希表+双向链表
		https://leetcode.cn/problems/lru-cache/solution/jian-dan-shi-li-xiang-xi-jiang-jie-lru-s-exsd/
	*/
}

type LRUCache struct {
	capacity int
	m        map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		capacity: capacity,
		m:        map[int]*Node{},
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) deleteNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) moveToHead(node *Node) {
	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) addToHead(node *Node) {
	this.head.Next.Prev = node
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next = node
}

func (this *LRUCache) removeTail() int {
	node := this.tail.Prev
	this.deleteNode(node)
	return node.Key
}

func (this *LRUCache) Get(key int) int {
	return -1
}

func (this *LRUCache) Put(key int, value int) {

}
