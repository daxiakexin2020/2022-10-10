package server

import (
	"22go_redis/server/construct"
	"errors"
	"time"
)

func (gr *Gredis) Set(key string, val interface{}, expirationTime time.Duration) {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	gr.Data[key] = construct.NewCString(val, expirationTime)
}

func (gr *Gredis) Get(key string) (interface{}, bool) {
	gr.MU.RLock()
	defer gr.MU.RUnlock()
	data, ok := gr.Data[key]
	if !ok {
		return nil, false
	}
	return data.GetVal(), true
}

func (gr *Gredis) STRLEN(key string) (int, error) {
	gr.MU.RLock()
	defer gr.MU.RUnlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return 0, errors.New("this key is not set")
	}
	if s, ok := vInterface.GetVal().(string); ok {
		return len(s), nil
	}
	return 0, errors.New("this is key is not string")
}

func (gr *Gredis) Decr(key string) error {
	gr.MU.Lock()
	defer gr.MU.Unlock()
	vInterface, ok := gr.Data[key]
	if !ok {
		return errors.New("this key is not set")
	}
	switch v := vInterface.GetVal().(type) {
	case int:
		vInterface.SetVal(v - 1)
		return nil
	default:
		return errors.New("this is key is not int, int8, int32, int64, uint, uint8, uint32, uint64....")
	}
}
