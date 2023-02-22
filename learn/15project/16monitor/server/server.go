package server

import (
	"context"
	"errors"
	"fmt"
)

type servicer interface {
	Name() string
	Monitor(ctx context.Context) error
	Alarm() error
}

type server struct {
	srvs []servicer
	boot bool
}

func NewServer() *server {
	return &server{srvs: make([]servicer, 0)}
}

func (s *server) Register(srv ...servicer) *server {
	s.srvs = append(s.srvs, srv...)
	return s
}

func (s *server) Run() error {
	if s.boot {
		return errors.New("server already starting.....")
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	for _, srv := range s.srvs {
		if err := srv.Monitor(ctx); err != nil {
			return err
		}
		fmt.Println(srv.Name() + " start ok....")
	}
	s.boot = true
	select {}
	return nil
}
