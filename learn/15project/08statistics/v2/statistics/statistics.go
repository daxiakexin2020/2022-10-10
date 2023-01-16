package statistics

import (
	"strings"
	"sync"
)

type Server struct {
	src       *service
	taskNames []string
	sl        sync.Mutex
	err       chan error
	isClosed  bool
}

var (
	sonce   sync.Once
	gserver *Server
)

const err_limit = 5

func NewServer() *Server {
	sonce.Do(func() {
		gserver = &Server{
			src: newService(),
			err: make(chan error, err_limit),
		}
	})
	return gserver
}

func (s *Server) Add(taskName string) {
	s.sl.Lock()
	defer s.sl.Unlock()
	if s.serIsClosed() {
		return
	}
	if err := s.src.add(taskName); err != nil {
		s.setError(err)
		return
	}
	s.taskNames = append(s.taskNames, taskName)
}

func (s *Server) end(taskName string) {
	if err := s.src.end(taskName); err != nil {
		s.setError(err)
	}
}

func (s *Server) Print(taskName string) (string, error) {
	s.sl.Lock()
	defer s.sl.Unlock()
	if s.serIsClosed() {
		return "服务已经关闭", nil
	}
	if len(s.err) >= 1 {
		err := <-s.err
		if err != nil {
			return "", err
		}
	}

	//调用End() 兜底补充结束参数
	s.end(taskName)

	//打印返回结果
	info, err := s.src.print(taskName)
	if err != nil {
		return "", err
	}
	return info, nil
}

func (s *Server) All() (string, error) {
	var str strings.Builder
	for _, taskname := range s.taskNames {
		res, err := s.Print(taskname)
		if err != nil {
			return str.String(), err
		}
		str.WriteString("\n")
		str.WriteString(res)
	}
	return str.String(), nil
}

func (s *Server) SetClosed() {
	s.setStatus(true)
}

func (s *Server) SetOpen() {
	s.setStatus(false)
}

func (s *Server) setStatus(status bool) {
	s.sl.Lock()
	defer s.sl.Unlock()
	s.isClosed = status
}

func (s *Server) serIsClosed() bool {
	return s.isClosed
}

func (s *Server) setError(err error) {
	if len(s.err) < err_limit {
		s.err <- err
	}
}
