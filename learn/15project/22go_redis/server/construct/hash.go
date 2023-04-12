package construct

type Chash struct {
	val map[string]interface{}
	*exTime
}

func NewChash() *Chash {
	return &Chash{val: map[string]interface{}{}}
}

func (ch *Chash) GetVal() interface{} {
	return ch.val
}

func (ch *Chash) Len() int {
	return len(ch.val)
}

func (ch *Chash) Del(field string) bool {
	if _, ok := ch.val[field]; ok {
		delete(ch.val, field)
		return ok
	}
	return false
}

func (ch *Chash) GetFieldVal(key string) interface{} {
	i, ok := ch.val[key]
	if !ok {
		return nil
	}
	return i
}

func (ch *Chash) SetVal(val interface{}) {}

func (ch *Chash) SetFieldVal(field string, val interface{}) bool {
	var exists bool
	if _, ok := ch.val[field]; !ok {
		exists = true
	}
	ch.val[field] = val
	return exists
}

func (ch *Chash) Type() string {
	return HASH
}
