package atomic

import "sync/atomic"

type Boolean uint32

func (b *Boolean) Get() bool {
	//判断一个uint32的数字是否是0，并发安全 原子操作
	return atomic.LoadUint32((*uint32)(b)) != 0
}

func (b *Boolean) Set(v bool) {
	if v {
		atomic.StoreUint32((*uint32)(b), 1)
	} else {
		atomic.StoreUint32((*uint32)(b), 0)
	}
}
