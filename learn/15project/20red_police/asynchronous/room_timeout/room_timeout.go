package room_timeout

import (
	"20red_police/asynchronous"
	"20red_police/asynchronous/room_timeout/lru"
	"sync"
	"time"
)

type roomTimeout struct {
	lru       *lru.ListNode
	mu        sync.Mutex
	boot      bool
	livetime  time.Duration
	OnEvicted func(roomID string) error
}

var (
	_  asynchronous.Tasker = (*roomTimeout)(nil)
	ec                     = make(chan struct{}, 1)
)
var (
	groomTimeout *roomTimeout
	ronce        sync.Once
)

func Timeout(capacity int64, livetime time.Duration, OnEvicted func(roomID string) error) *roomTimeout {
	ronce.Do(func() {
		groomTimeout = newRoomTimeout(capacity, livetime, OnEvicted)
	})
	return groomTimeout
}

func GTimeout() *roomTimeout {
	return groomTimeout
}

func newRoomTimeout(capacity int64, livetime time.Duration, OnEvicted func(roomID string) error) *roomTimeout {
	return &roomTimeout{
		lru:       lru.NewListNode(capacity),
		livetime:  livetime,
		OnEvicted: OnEvicted,
	}
}

func (rto *roomTimeout) TaskName() string {
	return "ROOM_TIMEOUT"
}

func (rto *roomTimeout) AddRoom(roomID string) {
	rto.mu.Lock()
	defer rto.mu.Unlock()
	rto.lru.Add(roomID)
}

func (rto *roomTimeout) ExitSignal() chan struct{} {
	return ec
}

func (rto *roomTimeout) Run() error {
	if rto.boot {
		return nil
	}
	rto.boot = true
	for {
		if !rto.boot {
			return nil
		}
		value := rto.lru.GetHeadValue(func(value string, addtime int64) bool {
			if time.Unix(addtime, 0).Add(rto.livetime).After(time.Now()) {
				return false
			}
			return true
		})
		if value != "" && rto.OnEvicted != nil {
			rto.OnEvicted(value)
		} else {
			time.Sleep(time.Second * 1)
		}
	}
	return nil
}

func (rto *roomTimeout) Stop() error {
	rto.boot = false
	return nil
}
