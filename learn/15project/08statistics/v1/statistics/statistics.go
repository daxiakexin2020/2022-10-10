package statistics

import (
	"strings"
	"sync"
)

type Server struct {
	taskNames []string
	timeSrc   *timeService
	memorySrc *memoryService
	sl        sync.Mutex
	err       chan error
}

var (
	sonce   sync.Once
	gserver *Server
)

const err_limit = 5

func NewServer() *Server {
	sonce.Do(func() {
		gserver = &Server{
			taskNames: make([]string, 0),
			timeSrc:   newTimeService(),
			memorySrc: newMemoryService(),
			err:       make(chan error, err_limit),
		}
	})
	return gserver
}

func (s *Server) Add(taskName string) {
	s.sl.Lock()
	defer s.sl.Unlock()
	if err := s.timeSrc.add(taskName); err != nil {
		s.setError(err)
		return
	}
	if err := s.memorySrc.add(taskName); err != nil {
		s.setError(err)
		return
	}
	s.taskNames = append(s.taskNames, taskName)
}

func (s *Server) end(taskName string) {
	if err := s.timeSrc.end(taskName); err != nil {
		s.setError(err)
	}
	if err := s.memorySrc.end(taskName); err != nil {
		s.setError(err)
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

	//调用End() 兜底补充结束参数
	s.end(taskName)

	//打印返回结果
	timeInfo, timeErr := s.timeSrc.print(taskName)
	memoryInfo, memoryErr := s.memorySrc.print(taskName)
	if timeErr != nil {
		return "", timeErr
	}
	if memoryErr != nil {
		return "", timeErr
	}
	return timeInfo + "============" + memoryInfo, nil
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
