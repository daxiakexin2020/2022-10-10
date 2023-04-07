package server

import (
	"22go_redis/server/construct"
	"time"
)

func (gr *Gredis) Ttl(key string) int64 {
	data, ok := gr.Data[key]
	if !ok {
		return -2
	}
	t := data.GetExTime()
	if t == 0 {
		return -1
	}
	return t - time.Now().Unix()
}

func (gr *Gredis) EXPIRE(key string, expirationTime time.Duration) bool {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return false
	}
	return vInterface.SetExTime(expirationTime)
}

func (gr *Gredis) Keys() []string {
	gr.MU.RLock()
	defer gr.MU.RUnlock()
	var keys []string
	for key, _ := range gr.Data {
		keys = append(keys, key)
	}
	return keys
}

func (gr *Gredis) Exists(key string) bool {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	_, ok := gr.Data[key]
	return ok
}

func (gr Gredis) Del(key string) int {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	if _, ok := gr.Data[key]; ok {
		delete(gr.Data, key)
		return 1
	}
	return -1
}

// get value type
func (gr *Gredis) Type(key string) string {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return ""
	}
	val := vInterface.GetVal()
	switch val.(type) {
	case construct.TypeString:
		return construct.STRING
	case construct.TypeSet:
		return construct.SET
	default:
		return construct.UNKNOWN
	}
}

func (gr *Gredis) RENAME(okey string, nkey string) bool {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[okey]
	if !ok {
		return false
	}
	gr.Data[nkey] = vInterface
	delete(gr.Data, okey)
	return true
}
