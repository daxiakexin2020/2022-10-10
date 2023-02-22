package main

import "fmt"

func main() {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	c.Put(3, 3)
	c.Put(1, 1)

	for k, v := range c.cache {
		fmt.Println(k, v)
	}
}

type LRUCache struct {
	size     int
	capacity int
	cache    map[int]*DLinkNode //key:DLinkNode
	head     *DLinkNode         //整个缓存系统的头节点
	tail     *DLinkNode         //整个缓存系统的尾节点
}

type DLinkNode struct {
	key   int
	value int
	prev  *DLinkNode
	next  *DLinkNode
}

func initDLinkNode(key int, value int) *DLinkNode {
	return &DLinkNode{key: key, value: value}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache:    map[int]*DLinkNode{},
		head:     initDLinkNode(0, 0),
		tail:     initDLinkNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {

	//新节点
	if _, ok := this.cache[key]; !ok {
		//创造节点
		node := initDLinkNode(key, value)

		//加入缓存
		this.cache[key] = node

		//加入头部
		this.addToHead(node)
		//计数+1
		this.size++
		//如果超出容量，删除尾部节点，删除缓存的节点
		if this.size > this.capacity {
			removeNode := this.removeTail()
			delete(this.cache, removeNode.key)
			this.size--
		}

		//旧节点，更新
	} else {
		node := this.cache[key]
		node.value = value
		//删除旧节点，并将其移动到头节点
		this.moveToHead(node)
	}
}

// 移除此节点
func (this *LRUCache) removeNode(node *DLinkNode) {

	//因为是双向链表，需要将此节点的上一个节点，下一个节点都重置一下指向

	//重置此节点的上一个节点的指向，将其下一个节点指向下下一个节点，跨过此节点
	node.prev.next = node.next

	//重置此节点的下一个节点的指向，将其指向此节点的上一个节点
	node.next.prev = node.prev

}

// 将节点移动到头部
func (this *LRUCache) moveToHead(node *DLinkNode) {

	//先移除此节点的信息
	this.removeNode(node)

	//再将其挂载在头部
	this.addToHead(node)
}

// 将节点添加在头部
func (this *LRUCache) addToHead(node *DLinkNode) {

	//主要分为2部分，一部分是为此节点配置上下游节点，另外一部分是为整体缓存重置head节点的上下指向

	//为此节点配置上一个节点，应该配置为当前的head节点
	node.prev = this.head

	//为此节点配置下一个节点，应该配置为当前头节点指向的节点
	node.next = this.head.next

	//
	this.head.next.prev = node

	//重置整体缓存的头节点的下一个节点的指向，指向要添加的节点
	this.head.next = node

}

func (this *LRUCache) removeTail() *DLinkNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}
