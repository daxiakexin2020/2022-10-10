package server

import (
	"22go_redis/server/construct"
	"22go_redis/utils"
)

func (gr *Gredis) SAdd(key string, val ...interface{}) int {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		uval := utils.SliceUniq(val)
		gr.Data[key] = construct.NewCset(uval, 0)
		return len(uval)
	} else {
		otv := vInterface.GetVal().([]interface{})
		tv := utils.SliceUniq(append(otv, val...))
		vInterface.SetVal(utils.SliceUniq(tv))
		return len(tv) - len(otv)
	}
}

func (gr *Gredis) Smembers(key string) []interface{} {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return nil
	}
	tv := vInterface.GetVal().([]interface{})
	return tv
}
