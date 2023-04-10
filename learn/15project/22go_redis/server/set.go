package server

import (
	"22go_redis/server/construct"
	"22go_redis/utils"
)

func (gr *Gredis) Sadd(key string, val ...interface{}) int {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	return gr.sadd(key, val...)
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

func (gr *Gredis) Scard(key string) int {
	gr.MU.RLock()
	defer gr.MU.RUnlock()
	if vInterface, ok := gr.Data[key]; !ok {
		return 0
	} else {
		if i, tok := vInterface.GetVal().([]interface{}); tok {
			return len(i)
		}
		return 0
	}
}

func (gr *Gredis) Sdiff(key string, keys ...string) []interface{} {
	gr.MU.RLock()
	gr.MU.RUnlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return nil
	}
	dest := vInterface.GetVal().([]interface{})
	var sdata []interface{}
	for _, k := range keys {
		if v, ok := gr.Data[k]; ok {
			if i, iok := v.GetVal().([]interface{}); iok {
				sdata = append(sdata, i...)
			}
		}
	}
	return utils.SliceDiff(dest, sdata)
}

func (gr *Gredis) Smove(key string, destKey string, m interface{}) int {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return 0
	}
	source := vInterface.GetVal().([]interface{})
	if !utils.InSlice(source, m) {
		return 0
	}

	dvInterface, ok := gr.Data[destKey]
	if !ok {
		gr.sadd(destKey, m)
	} else {
		destSource := dvInterface.GetVal().([]interface{})
		if utils.InSlice(destSource, m) {
			return 0
		}
		destSource = append(destSource, m)
		dvInterface.SetVal(m)
	}

	for i, data := range source {
		if data == m {
			source = append(source[0:i], source[i+1:]...)
			vInterface.SetVal(source)
			break
		}
	}
	return 1
}

func (gr *Gredis) Spop(key string, count int) []interface{} {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return nil
	}
	dest := vInterface.GetVal().([]interface{})
	if count >= len(dest) {
		return dest
	}

	indexs := utils.RandInts(len(dest), count)
	if len(indexs) == 0 {
		return nil
	}

	var ret []interface{}
	var restDest []interface{}
	flag := make(map[interface{}]struct{})
	for _, i := range indexs {
		ret = append(ret, dest[i])
		flag[dest[i]] = struct{}{}
	}
	for _, v := range dest {
		if _, ok := flag[v]; !ok {
			delete(flag, v)
			restDest = append(restDest, v)
		}
	}
	vInterface.SetVal(restDest)
	return ret
}

func (gr *Gredis) Srandmember(key string, count int) []interface{} {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return nil
	}
	dest := vInterface.GetVal().([]interface{})
	if count >= len(dest) {
		return dest
	}

	indexs := utils.RandInts(len(dest), count)
	if len(indexs) == 0 {
		return nil
	}
	var ret []interface{}
	for _, i := range indexs {
		ret = append(ret, dest[i])
	}
	return ret
}

func (gr *Gredis) Srem(key string, m ...interface{}) int {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return 0
	}
	flag := make(map[interface{}]struct{})
	for _, v := range m {
		flag[v] = struct{}{}
	}
	var resetDest []interface{}
	var ret int
	dest := vInterface.GetVal().([]interface{})
	for _, v := range dest {
		if _, ok := flag[v]; ok {
			ret++
		} else {
			resetDest = append(resetDest, v)
		}
	}
	vInterface.SetVal(resetDest)
	return ret
}

func (gr *Gredis) sadd(key string, val ...interface{}) int {
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
