package construct

import (
	"sync"
	"time"
)

type VInterface interface {
	Type() string
	GetVal() interface{}
	SetVal(val interface{})
	GetExTime() int64
	SetExTime(t time.Duration) bool
}

type CGORedis struct {
	data map[string]VInterface
	mu   sync.RWMutex
}

func NewCGORedis() *CGORedis {
	return &CGORedis{
		data: map[string]VInterface{},
	}
}

func (cgo *CGORedis) CGOSet(key string, val VInterface) {
	cgo.data[key] = val
}

func (cgo *CGORedis) CGOGet(key string) VInterface {
	return cgo.data[key]
}

func (cgo *CGORedis) CGODel(key string) {
	delete(cgo.data, key)
}

func (cgo *CGORedis) CGOIsData(key string) (VInterface, bool) {
	if vInterface, ok := cgo.data[key]; ok {
		return vInterface, ok
	}
	return nil, false
}

func (cgo *CGORedis) CGOKeys() []string {
	var keys []string
	for k, _ := range cgo.data {
		keys = append(keys, k)
	}
	return keys
}

func (cgo *CGORedis) Lock() {
	cgo.mu.Lock()
}

func (cgo *CGORedis) Unlock() {
	cgo.mu.Unlock()
}

func (cgo *CGORedis) RLock() {
	cgo.mu.RLock()
}

func (cgo *CGORedis) RUnlock() {
	cgo.mu.RUnlock()
}

const (
	STRING  = "string"
	LIST    = "list"
	SET     = "set"
	ZSET    = "zset"
	HASH    = "hash"
	BITMAP  = "bitmap"
	UNKNOWN = "unknown"
)
