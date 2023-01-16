package statistics

import (
	"strings"
	"sync"
)

type Server struct {
	taskNames []string
	srcs      []servicer
	sl        sync.Mutex
	err       chan error
}

type servicer interface {
	add(taskName string) error
	end(taskName string) error
	print(taskName string) (string, error)
}

var (
	sonce   sync.Once
	gserver *Server
)

const err_limit = 5

func NewServer() *Server {
	sonce.Do(func() {
		srcs := make([]servicer, 0)
		srcs = append(srcs, newTimeService())
		srcs = append(srcs, newMemoryService())
		gserver = &Server{
			taskNames: make([]string, 0),
			srcs:      srcs,
			err:       make(chan error, err_limit),
		}
	})
	return gserver
}

func (s *Server) Add(taskName string) {
	s.sl.Lock()
	defer s.sl.Unlock()
	for _, task := range s.srcs {
		if err := task.add(taskName); err != nil {
			s.setError(err)
			return
		}
	}
	s.taskNames = append(s.taskNames, taskName)
}

func (s *Server) end(taskName string) {
	for _, task := range s.srcs {
		if err := task.end(taskName); err != nil {
			s.setError(err)
		}
	}
}

func (s *Server) Print(taskName string) (string, error) {
	s.sl.Lock()
	defer s.sl.Unlock()

	if len(s.err) >= 1 {
		err := <-s.err
		if err != nil {
			return "", err
		}
	}

	//调用end() 兜底补充结束参数
	s.end(taskName)

	//打印返回结果
	var str strings.Builder
	for _, task := range s.srcs {
		strInfo, err := task.print(taskName)
		if err != nil {
			return "", err
		}
		str.WriteString("================")
		str.WriteString(strInfo)
	}
	return str.String(), nil
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

func (s *Server) setError(err error) {
	if len(s.err) < err_limit {
		s.err <- err
	}
}
