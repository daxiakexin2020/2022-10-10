package construct

import "time"

type Cset struct {
	val []interface{}
	*exTime
}

var (
	_ VInterface = (*Cset)(nil)
)

func NewCset(val []interface{}, etime time.Duration) *Cset {
	return &Cset{
		val:    val,
		exTime: NewExTime(etime),
	}
}

func (cs *Cset) Type() string {
	return SET
}

func (cs *Cset) GetVal() interface{} {
	return cs.val
}

func (cs *Cset) SetVal(val interface{}) {
	if i, ok := val.([]interface{}); ok {
		cs.val = i
	}
}

func (cs *Cset) Len() int {
	return len(cs.val)
}
