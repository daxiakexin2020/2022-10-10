package main

import "sync"

type MyHashSet struct {
	keys map[int]Empty
	lock sync.Mutex
}

type Empty struct{}

/*
*
不使用任何内建的哈希表库设计一个哈希集合（HashSet）。

实现 MyHashSet 类：

void add(key) 向哈希集合中插入值 key 。
bool contains(key) 返回哈希集合中是否存在这个值 key 。
void remove(key) 将给定值 key 从哈希集合中删除。如果哈希集合中没有这个值，什么也不做。
*/
func main() {

}

func Constructor() MyHashSet {
	return MyHashSet{
		keys: map[int]Empty{},
		lock: sync.Mutex{},
	}
}

func (this *MyHashSet) Add(key int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.keys[key] = Empty{}
}

func (this *MyHashSet) Remove(key int) {
	this.lock.Lock()
	defer this.lock.Unlock()
	if _, ok := this.keys[key]; ok {
		delete(this.keys, key)
	}
}

func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.keys[key]
	return ok
}
