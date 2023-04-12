package items

import (
	"22go_redis/server/construct"
	cerror "22go_redis/server/error"
)

func (gr *Gredis) Hset(key string, field string, val interface{}) (bool, error) {
	gr.Lock()
	defer gr.Unlock()
	data, b := gr.isData(key)
	chash := construct.NewChash()
	var set bool
	if !b {
		set = chash.SetFieldVal(field, val)
	} else {
		if data.Type() != construct.HASH {
			return false, cerror.HASH_TYPE_ERROR
		}
		chash = data.(*construct.Chash)
		set = chash.SetFieldVal(field, val)
	}
	gr.CGOSet(key, chash)
	return set, nil
}

func (gr *Gredis) Hlen(key string) int {
	gr.RLock()
	defer gr.RUnlock()
	data, b := gr.isData(key)
	if !b {
		return 0
	}
	if data.Type() != construct.HASH {
		return 0
	}
	chash := data.(*construct.Chash)
	return chash.Len()
}

func (gr *Gredis) Hget(key string, field string) interface{} {
	gr.Lock()
	defer gr.Unlock()
	data, b := gr.isData(key)
	if !b {
		return nil
	}
	if data.Type() != construct.HASH {
		return nil
	}
	return data.(*construct.Chash).GetFieldVal(field)
}

func (gr *Gredis) Hdel(key string, fields ...string) int {
	gr.Lock()
	defer gr.Unlock()
	data, b := gr.isData(key)
	if !b {
		return 0
	}
	if data.Type() != construct.HASH {
		return 0
	}
	var count int
	chash := data.(*construct.Chash)
	for _, field := range fields {
		del := chash.Del(field)
		if del {
			count++
		}
	}
	return count
}
