package main

import (
	"fmt"
	"time"
)

/*
*
实现 LFUCache 类：

LFUCache(int capacity) - 用数据结构的容量 capacity 初始化对象
int get(int key) - 如果键 key 存在于缓存中，则获取键的值，否则返回 -1 。
void put(int key, int value) - 如果键 key 已存在，则变更其值；如果键不存在，请插入键值对。
当缓存达到其容量 capacity 时，则应该在插入新项之前，移除最不经常使用的项。在此问题中，当存在平局（即两个或更多个键具有相同使用频率）时，应该去除 最近最久未使用 的键。
为了确定最不常使用的键，可以为缓存中的每个键维护一个 使用计数器 。使用计数最小的键是最久未使用的键。

当一个键首次插入到缓存中时，它的使用计数器被设置为 1 (由于 put 操作)。对缓存中的键执行 get 或 put 操作，使用计数器的值将会递增。

函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
*/
func main() {

	//["LFUCache","get","put","get","put","put","get","get"]
	//[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]
	l := Constructor(2)

	l.Get(2)
	l.Put(2, 6)
	l.Get(1)

	l.Put(1, 5)
	l.Put(1, 2)

	fmt.Println("1:", l.Get(1)) //2   2
	fmt.Println("2:", l.Get(2)) //6  -1

	return

	fmt.Println(l.datas) //1,3  没有问题

	l.Get(2)    //-1
	l.Get(3)    // 3
	l.Put(4, 4) //4 3

	fmt.Println(l.datas) //1 3这里没有将1挪走，而是把4挪走了///////

	fmt.Println("1:", l.Get(1)) //-1
	fmt.Println("3:", l.Get(3)) //3
	fmt.Println("4:", l.Get(4)) //4

	fmt.Println(l.datas)
}

type LFUCache struct {
	cap   int
	datas map[int]int
	flags map[int]entry
}

type entry struct {
	count int
	date  int64
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:   capacity,
		datas: make(map[int]int),
		flags: make(map[int]entry),
	}
}

func (this *LFUCache) Get(key int) int {
	if data, ok := this.datas[key]; ok {
		e := this.flags[key]
		e.count += 1
		e.date = time.Now().UnixNano()
		this.flags[key] = e
		return data
	}
	return -1
}

var countflag = 1000000
var dataflag int64
var flagKey int

func (this *LFUCache) Put(key int, value int) {

	_, ok := this.datas[key]

	if len(this.datas) == this.cap && !ok {
		for k, _ := range this.datas {
			c := this.flags[k].count
			d := this.flags[k].date
			if c > countflag {
				continue
			}
			//找到更小的了,将标识数重置，将keys清空，加入新的key
			if c < countflag {
				countflag = c
				dataflag = d
				flagKey = k
			} else {
				//相等，需要比较日期
				if d <= dataflag {
					countflag = c
					dataflag = d
					flagKey = k
				}
			}
		}
		delete(this.datas, flagKey)
		delete(this.flags, flagKey)
	}
	this.datas[key] = value
	e := this.flags[key]
	e.count += 1
	e.date = time.Now().UnixNano()
	this.flags[key] = e

	countflag = 1000000
	dataflag = 0
	flagKey = 0
}
