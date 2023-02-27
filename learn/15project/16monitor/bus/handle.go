package bus

import (
	"16monitor/server"
	"16monitor/server/kinds"
)

func Handle() {
	server := server.NewServer()
	server.Register(
		kinds.NewMemory(98, 3),
		kinds.NewDF(1, 1))
	server.Run()
}
