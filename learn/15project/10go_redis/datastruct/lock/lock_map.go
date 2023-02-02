package lock

import (
	"sort"
	"sync"
)

/**
../dict/concurrent.go 中 实现的ConcurrentMap 可以保证对单个 key 操作的并发安全性，但是仍然无法满足需求:
Incr 命令需要完成: 读取 -> 做加法 -> 写入 三步操作
MSETNX 命令当且仅当所有给定键都不存在时所有给定键设置值, 因此我们需要锁定所有给定的键直到完成所有键的检查和设置
因此我们需要实现 db.Locker 用于锁定一个或一组 key 并在我们需要的时候释放锁。
实现 db.Locker 最直接的想法是使用一个 map[string]*sync.RWMutex, 加锁过程分为两步: 初始化对应的锁 -> 加锁， 解锁过程也分为两步: 解锁 -> 释放对应的锁。那么存在一个无法解决的并发问题
*/

const (
	prime32 = uint32(16777619)
)

type Locks struct {
	table []*sync.RWMutex
}

func Make(tableSize int) *Locks {
	table := make([]*sync.RWMutex, tableSize)
	for i := 0; i < tableSize; i++ {
		table[i] = &sync.RWMutex{}
	}
	return &Locks{
		table: table,
	}
}

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

func (locks *Locks) spread(hashCode uint32) uint32 {
	if locks == nil {
		panic("locks is nil")
	}
	tableSize := uint32(len(locks.table))
	return (tableSize - 1) & uint32(hashCode)
}

func (locks *Locks) Lock(key string) {
	index := locks.spread(fnv32(key))
	mu := locks.table[index]
	mu.Lock()
}

func (Locks *Locks) RLock(key string) {
	index := Locks.spread(fnv32(key))
	mu := Locks.table[index]
	mu.RLock()
}

func (Locks *Locks) UnLock(key string) {
	index := Locks.spread(fnv32(key))
	mu := Locks.table[index]
	mu.Unlock()
}

func (Locks *Locks) RUnLock(key string) {
	index := Locks.spread(fnv32(key))
	mu := Locks.table[index]
	mu.RUnlock()
}

/**
在锁定多个key时需要注意，若协程A持有键a的锁试图获得键b的锁，此时协程B持有键b的锁试图获得键a的锁则会形成死锁。
解决方法是所有协程都按照相同顺序加锁，若两个协程都想获得键a和键b的锁，那么必须先获取键a的锁后获取键b的锁，这样就可以避免循环等待。
*/

func (locks *Locks) toLockIndices(keys []string, reverse bool) []uint32 {
	indexs := make(map[uint32]bool)
	for _, key := range keys {
		//获取单key对应的index，在map中作标记
		index := locks.spread(fnv32(key))
		indexs[index] = true
	}

	//把index放在切片中，方便排序
	indices := make([]uint32, 0, len(indexs))
	for index := range indexs {
		indices = append(indices, index)
	}

	//将index按照大》小或者小》大的顺序排序
	sort.Slice(indices, func(i, j int) bool {
		if !reverse {
			return indices[i] < indices[j]
		}
		return indices[i] > indices[j]
	})
	return indices
}

func (locks *Locks) Locks(keys ...string) {
	indices := locks.toLockIndices(keys, false)
	for _, index := range indices {
		mu := locks.table[index]
		mu.Lock()
	}
}
func (locks *Locks) RLocks(keys ...string) {
	indices := locks.toLockIndices(keys, false)
	for _, index := range indices {
		mu := locks.table[index]
		mu.RLock()
	}
}

// UnLocks releases multiple exclusive locks
func (locks *Locks) UnLocks(keys ...string) {
	indices := locks.toLockIndices(keys, true)
	for _, index := range indices {
		mu := locks.table[index]
		mu.Unlock()
	}
}

// RUnLocks releases multiple shared locks
func (locks *Locks) RUnLocks(keys ...string) {
	indices := locks.toLockIndices(keys, true)
	for _, index := range indices {
		mu := locks.table[index]
		mu.RUnlock()
	}
}

// RWLocks locks write keys and read keys together. allow duplicate keys
func (locks *Locks) RWLocks(writeKeys []string, readKeys []string) {
	keys := append(writeKeys, readKeys...)
	indices := locks.toLockIndices(keys, false)
	writeIndexSet := make(map[uint32]struct{})
	for _, wKey := range writeKeys {
		idx := locks.spread(fnv32(wKey))
		writeIndexSet[idx] = struct{}{}
	}
	for _, index := range indices {
		_, w := writeIndexSet[index]
		mu := locks.table[index]
		if w {
			mu.Lock()
		} else {
			mu.RLock()
		}
	}
}

// RWUnLocks unlocks write keys and read keys together. allow duplicate keys
func (locks *Locks) RWUnLocks(writeKeys []string, readKeys []string) {
	keys := append(writeKeys, readKeys...)
	indices := locks.toLockIndices(keys, true)
	writeIndexSet := make(map[uint32]struct{})
	for _, wKey := range writeKeys {
		idx := locks.spread(fnv32(wKey))
		writeIndexSet[idx] = struct{}{}
	}
	for _, index := range indices {
		_, w := writeIndexSet[index]
		mu := locks.table[index]
		if w {
			mu.Unlock()
		} else {
			mu.RUnlock()
		}
	}
}
