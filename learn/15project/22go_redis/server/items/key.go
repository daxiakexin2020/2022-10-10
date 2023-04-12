package items

import (
	"22go_redis/server/construct"
	"time"
)

func (gr *Gredis) Ttl(key string) int64 {

	isData, b := gr.isData(key)
	if !b {
		return -2
	}
	t := isData.GetExTime()
	if t == 0 {
		return -1
	}
	return t - time.Now().Unix()
}

func (gr *Gredis) EXPIRE(key string, expirationTime time.Duration) bool {
	gr.Lock()
	defer gr.Unlock()
	vInterface, ok := gr.isData(key)
	if !ok {
		return false
	}
	return vInterface.SetExTime(expirationTime)
}

func (gr *Gredis) Keys() []string {
	gr.RLock()
	defer gr.RUnlock()
	return gr.CGOKeys()
}

func (gr *Gredis) Exists(key string) bool {
	gr.Lock()
	defer gr.Unlock()
	_, ok := gr.isData(key)
	return ok
}

func (gr Gredis) Del(key string) int {
	gr.Lock()
	defer gr.Unlock()
	if _, ok := gr.isData(key); ok {
		gr.CGODel(key)
		return 1
	}
	return -1
}

// get value type
func (gr *Gredis) Type(key string) string {
	gr.RLock()
	defer gr.RUnlock()
	vInterface, ok := gr.isData(key)
	if !ok {
		return construct.UNKNOWN
	}
	switch vInterface.(type) {
	case *construct.Cstring:
		return construct.STRING
	case *construct.Cset:
		return construct.SET
	case *construct.CZset:
		return construct.ZSET
	case *construct.Clist:
		return construct.LIST
	case *construct.Chash:
		return construct.HASH
	default:
		return construct.UNKNOWN
	}
}

func (gr *Gredis) RENAME(okey string, nkey string) bool {
	gr.Lock()
	defer gr.Unlock()
	data, b := gr.isData(okey)
	if !b {
		return false
	}
	gr.CGOSet(nkey, data)
	gr.Del(okey)
	return true
}
