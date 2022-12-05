package start_center

import (
	"errors"
	"sync"
)

type Handle func() error

type Server struct {
	H      []Handle
	isBoot bool
	lock   sync.Mutex
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Register(handles ...Handle) {
	s.H = append(s.H, handles...)
}

func (s *Server) Run() error {
	if s.isBoot {
		return errors.New("已经启动服务")
	}
	return s.run()
}

func (s *Server) run() error {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, handle := range s.H {
		if err := handle(); err != nil {
			return err
		}
	}
	s.isBoot = true
	return nil
}
