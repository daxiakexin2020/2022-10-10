package items

import "22go_redis/server/construct"

func (gr *Gredis) Llen(key string) int32 {
	gr.RLock()
	defer gr.RUnlock()
	data, b := gr.isData(key)
	if !b {
		return 0
	}
	if data.Type() != construct.LIST {
		return 0
	}
	clist := data.(*construct.Clist)
	return clist.Size()
}

func (gr *Gredis) Lpush(key string, val ...interface{}) {
	gr.Lock()
	defer gr.Unlock()
	data, b := gr.isData(key)
	clist := &construct.Clist{}
	if !b {
		clist = construct.NewClist(val[0])
		val = val[1:]
	} else {
		clist = data.(*construct.Clist)
	}
	for _, v := range val {
		clist.AddToHead(v)
	}
	gr.CGOSet(key, clist)
}

func (gr *Gredis) Lpop(key string) interface{} {
	gr.Lock()
	defer gr.Unlock()
	data, b := gr.isData(key)
	if !b {
		return nil
	}
	c := data.(*construct.Clist)
	head := c.RemoveHead()
	if head == nil {
		return nil
	}
	return head.GetVal()
}

func (gr *Gredis) Lindex(key string, index int) interface{} {
	gr.Lock()
	gr.Unlock()
	data, b := gr.isData(key)
	if !b {
		return nil
	}
	clist := data.(*construct.Clist)
	node := clist.Get(index)
	if node == nil {
		return nil
	}
	return node.GetVal()
}

func (gr *Gredis) LpushX(key string, val interface{}) {
	gr.Lock()
	defer gr.Unlock()
	data, b := gr.isData(key)
	if !b {
		return
	}
	clist := data.(*construct.Clist)
	clist.AddToHead(val)
	gr.CGOSet(key, clist)
}

func (gr *Gredis) Lrange(key string, start int, end int) []interface{} {
	gr.Lock()
	defer gr.Unlock()
	data, b := gr.isData(key)
	if !b {
		return nil
	}
	clist := data.(*construct.Clist)
	nodes := clist.GetRange(start, end)
	var ret []interface{}
	if nodes == nil {
		return ret
	}
	for _, node := range nodes {
		ret = append(ret, node.GetVal())
	}
	return ret
}

func (gr *Gredis) Lrem(key string, count int32, val interface{}) int32 {
	gr.Lock()
	defer gr.Unlock()
	data, b := gr.isData(key)
	if !b {
		return 0
	}
	clist := data.(*construct.Clist)
	return clist.RemoveVal(count, val)
}

func (gr *Gredis) Lset(key string, index int, val interface{}) bool {
	gr.Lock()
	defer gr.Unlock()
	data, b := gr.isData(key)
	if !b {
		return false
	}
	if data.Type() != construct.LIST {
		return false
	}
	clists := data.(*construct.Clist)
	node := clists.Get(index)
	if node == nil {
		return false
	}
	node.SetVal(val)
	gr.CGOSet(key, clists)
	return true
}
