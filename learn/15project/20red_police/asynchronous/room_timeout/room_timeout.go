package room_timeout

import (
	"20red_police/asynchronous"
	"20red_police/asynchronous/room_timeout/lru"
	"fmt"
	"sync"
)

type RoomTimeOut struct {
	lru  *lru.ListNode
	mu   sync.Mutex
	boot bool
}

var _ asynchronous.Tasker = (*RoomTimeOut)(nil)
var es = make(chan struct{}, 1)

func NewRoomTimeOut(capacity int64, OnEvicted chan string) asynchronous.Tasker {
	return &RoomTimeOut{lru: lru.NewListNode(capacity, OnEvicted)}
}

func (rto *RoomTimeOut) TaskName() string {
	return "ROOM_TIMEOUT"
}

func (rto *RoomTimeOut) AddRoom(roomID string) {
	rto.mu.Lock()
	defer rto.mu.Unlock()
	rto.lru.Add(roomID)
}

func (rto *RoomTimeOut) ExitSignal() chan struct{} {
	return es
}

func (rto *RoomTimeOut) Run() error {
	if rto.boot {
		return nil
	}
	rto.boot = true
	for {
		if !rto.boot {
			return nil
		}
		value := rto.lru.GetHeadValue(true)
		if value != "" {
			fmt.Println("run node value:", value)
			//业务
			if rto.lru.OnEvicted != nil {
			}
		}
	}
	return nil
}

func (rto *RoomTimeOut) Stop() error {
	rto.boot = false
	return nil
}
