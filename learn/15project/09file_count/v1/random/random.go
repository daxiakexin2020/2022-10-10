package random

import "sync"

// 100w+数据，  top10
// 让6, 66, 666, 6666, 66666, 8, 88, 888, 8888, 88888 最多
type Rand struct {
	su *shareUnit
}

type empty struct{}

// 享元
type shareUnit struct {
	flags map[int]empty
}

var (
	su     *shareUnit
	suOnce sync.Once
)

var flags = []int{
	6, 66, 666, 6666, 66666,
	8, 88, 888, 8888, 88888,
}

func makeShareUnit() *shareUnit {
	suOnce.Do(func() {
		f := make(map[int]empty)
		for _, flag := range flags {
			f[flag] = empty{}
		}
		su = &shareUnit{
			flags: f,
		}
	})
	return su
}

func NewRand() *Rand {
	return &Rand{su: makeShareUnit()}
}

// 返回此数字，需要生成多少次
func (r *Rand) MakeCount(dest int32) int32 {
	if _, ok := r.su.flags[int(dest)]; ok {
		return dest
	}
	return 1
}
