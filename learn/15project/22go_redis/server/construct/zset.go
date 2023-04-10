package construct

import (
	"time"
)

type CZset struct {
	val []*Czval
	*exTime
}

type Czval struct {
	val   interface{}
	score int64
}

var (
	_ VInterface = (*CZset)(nil)
)

func NewCzval(val interface{}, score int64) *Czval {
	return &Czval{val: val, score: score}
}

func NewCZset(etime time.Duration) *CZset {
	return &CZset{
		val:    make([]*Czval, 0),
		exTime: NewExTime(etime),
	}
}

func (cz *CZset) Type() string {
	return "zset"
}

func (cz *CZset) GetVal() interface{} {
	return cz.val
}

func (cz *CZset) SetVal(val interface{}) {
	if i, ok := val.([]*Czval); ok {
		cz.val = i
	}
}

func (czv *Czval) GetVal() interface{} {
	return czv.val
}

func (czv *Czval) GetScore() int64 {
	return czv.score
}
