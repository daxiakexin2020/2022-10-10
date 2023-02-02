package dict

import (
	"math"
	"sync"
	"sync/atomic"
)

type ConcurrentDict struct {
	table      []*shard
	count      int32
	shardCount int
}

/*
*
采用 golang 社区广泛使用的分段锁策略。
我们将 key 分散到固定数量的 shard 中避免 rehash 操作。shard 是有锁保护的 map, 当 shard 进行 rehash 时会阻塞shard内的读写，但不会对其他 shard 造成影响。
*/
type shard struct {
	m     map[string]interface{}
	mutex sync.RWMutex
}

func computeCapacity(param int) (size int) {
	if param <= 16 {
		return 16
	}
	n := param - 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	if n < 0 {
		return math.MaxInt32
	}
	return n + 1
}

func MakeConcurrent(shardCount int) *ConcurrentDict {
	shardCount = computeCapacity(shardCount)
	table := make([]*shard, shardCount)
	for j := 0; j < shardCount; j++ {
		table[j] = &shard{
			m: make(map[string]interface{}),
		}
	}
	return &ConcurrentDict{
		count:      0,
		table:      table,
		shardCount: shardCount,
	}
}

// hash fnv算法
const prime32 = uint32(16777619)

// 传入一个key，返回一个uint32,hash值
func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// 定位shard, 当n为2的整数幂时 h % n == (n - 1) & h
func (dict *ConcurrentDict) spread(hashCode uint32) uint32 {
	if dict == nil {
		panic("dict is nil")
	}
	tableSize := uint32(len(dict.table))
	/**
	^ 按位异或  当二进制位有一个为1，另外一个0时，为1，否则为0
	| 按位或    当二进制位只要有一个为1，则为1，否则为0
	& 按位与，当二进制位同时为1时，为1，否则为0
	*/
	return (tableSize - 1) & hashCode
}

func (dict *ConcurrentDict) getShard(index uint32) *shard {
	if dict == nil {
		panic("dict is nil")
	}
	return dict.table[index]
}

func (dict *ConcurrentDict) Get(key string) (val interface{}, exists bool) {
	if dict == nil {
		panic("dict is nil")
	}
	//先拿hashCode
	hashCode := fnv32(key)

	//再拿对应的shard的index
	index := dict.spread(hashCode)

	//拿到shard
	s := dict.getShard(index)

	//上锁，获取map值
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	val, exists = s.m[key]
	return
}

func (dict *ConcurrentDict) Put(key string, value interface{}) (result int) {
	if dict == nil {
		panic("dict is nil")
	}
	//先计算此kay的hashCode
	hashCode := fnv32(key)
	//拿index
	index := dict.spread(hashCode)
	//拿shard
	s := dict.getShard(index)
	//上锁开始操作
	s.mutex.Lock()
	defer s.mutex.Unlock()

	//1 此key在map中，覆盖旧值，返回
	if _, ok := s.m[key]; ok {
		s.m[key] = value
		return 0
	}
	//2 此key不在map中，计数+1
	s.m[key] = value
	dict.addCount()
	return 1
}

func (dict *ConcurrentDict) PutIfAbsent(key string, val interface{}) (result int) {
	if dict == nil {
		panic("dict is nil")
	}
	hashCode := fnv32(key)
	index := dict.spread(hashCode)
	s := dict.getShard(index)
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if _, ok := s.m[key]; ok {
		return 0
	}
	s.m[key] = val
	dict.addCount()
	return 1
}

func (dict *ConcurrentDict) addCount() int32 {
	return atomic.AddInt32(&dict.count, 1)
}

func (dict *ConcurrentDict) decreaseCount() int32 {
	return atomic.AddInt32(&dict.count, -1)
}

func (dict *ConcurrentDict) Len() int {
	if dict == nil {
		panic("dict is nil")
	}
	return int(atomic.LoadInt32(&dict.count))
}
