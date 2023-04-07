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
	Data map[string]VInterface
	MU   sync.RWMutex
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

type (
	TypeString interface{}
	TypeSet    []interface{}
)

func NewCGORedis() *CGORedis {
	return &CGORedis{
		Data: map[string]VInterface{},
	}
}

type exTime struct {
	time int64
}

func (ex *exTime) GetExTime() int64 {
	return ex.time
}

func (ex *exTime) SetExTime(t time.Duration) bool {
	ex.time = time.Now().Add(time.Second * t).Unix()
	return true
}

func NewExTime(t time.Duration) *exTime {
	return &exTime{time: time.Now().Add(time.Second * t).Unix()}
}
