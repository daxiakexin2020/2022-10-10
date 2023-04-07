package construct

import "time"

type Cstring struct {
	val interface{}
	*exTime
}

var (
	_ VInterface = (*Cstring)(nil)
)

func NewCString(val interface{}, etime time.Duration) *Cstring {
	return &Cstring{
		val:    val,
		exTime: NewExTime(etime),
	}
}

func (cst *Cstring) Type() string {
	return "string"
}

func (cst *Cstring) GetVal() interface{} {
	return cst.val
}

func (cst *Cstring) SetVal(val interface{}) {
	cst.val = val
}
