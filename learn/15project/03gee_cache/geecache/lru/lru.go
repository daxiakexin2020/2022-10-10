package lru

import "container/list"

/**
创建包含字典和双向链表的结构体类型 Cache，方便实现后续的增删查改操作。
  	在这里我们直接使用 Go 语言标准库实现的双向链表list.List。
	字典的定义是 map[string]*list.Element，键是字符串，值是双向链表中对应节点的指针。
	maxBytes 是允许使用的最大内存，nbytes 是当前已使用的内存，OnEvicted 是某条记录被移除时的回调函数，可以为 nil。
	键值对 entry 是双向链表节点的数据类型，在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
	为了通用性，我们允许值是实现了 Value 接口的任意类型，该接口只包含了一个方法 Len() int，用于返回值所占用的内存大小。
*/

type Cache struct {
	maxBytes  int64                         //是允许使用的最大内存
	nbytes    int64                         //是当前已使用的内存
	ll        *list.List                    //双向链表
	cache     map[string]*list.Element      //键是字符串，值是双向链表中对应节点的指针。
	OnEvicted func(key string, value Value) //是某条记录被移除时的回调函数，可以为 nil
}

// 键值对 entry 是双向链表节点的数据类型，在链表中仍保存每个值对应的 key 的好处在于，淘汰队首节点时，需要用 key 从字典中删除对应的映射。
type entry struct {
	key   string
	value Value
}

// 为了通用性，我们允许值是实现了 Value 接口的任意类型，该接口只包含了一个方法 Len() int，用于返回值所占用的内存大小。
type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)    //将此元素移动到双向链表队尾 双向链表作为队列，队首队尾是相对的，在这里约定 front 为队尾
		kv := ele.Value.(*entry) //断言成entry
		return kv.value, true    //返回value
	}
	return
}

// 这里的删除，实际上是缓存淘汰。即移除最近最少访问的节点（队首）
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back() //取到队首节点
	if ele != nil {
		c.ll.Remove(ele) //从链表中删除
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)                                //从字典中，删除该节点的映射关系
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len()) //当前使用的缓存，减去此节点
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 新增/修改
func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(len(key)) - int64(value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key: key, value: value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes > 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
