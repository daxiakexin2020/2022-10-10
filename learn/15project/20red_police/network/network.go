package network

import (
	"net"
	"sync"
)

type Resources struct {
	conns map[string]net.Conn
	mu    sync.Mutex
}

var (
	gresoureces *Resources
)

func init() {
	gresoureces = &Resources{conns: map[string]net.Conn{}}
}

func Gresources() *Resources {
	return gresoureces
}

func (r *Resources) add(threadID string, conn net.Conn) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.conns[threadID] = conn
}

func (r *Resources) Broadcast(data interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, conn := range r.conns {
		go sendResponse(data, nil, conn)
	}
}
