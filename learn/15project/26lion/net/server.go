package net

import (
	"26lion/config"
	"26lion/iface"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync/atomic"
)

type Server struct {
	Name      string
	IPVersion string
	Ip        string
	Port      int
	exitCh    chan struct{}
}

type Option func(s *Server)

var LionServer = NewServer()

var _ iface.IServer = (*Server)(nil)

func NewServer(opts ...Option) *Server {
	s := &Server{
		Name:      config.GetServerConfig().Name,
		IPVersion: config.GetServerConfig().IPVersion,
		Ip:        config.GetServerConfig().Ip,
		Port:      config.GetServerConfig().Port,
		exitCh:    make(chan struct{}),
	}
	s.apply(opts...)
	return s
}

func (s *Server) Serve() {
	s.Start()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Printf("[SERVE] Lion server , name %s, Serve Interrupt, signal = %v", s.Name, sig)
}

func (s *Server) Start() {
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf(":%s:%d", s.Ip, s.Port))
		if err != nil {
			log.Println("Net.ResolveTCPAddr err:", err)
			return
		}

		listen, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			log.Println("Net.ListenTCP err:", err)
			return
		}

		var connId uint64

		go func() {
			for {
				conn, err := listen.Accept()
				if err != nil {
					log.Println("Listen.Accept err:", err)
					continue
				}
				lconn := newConnection(s, conn, connId)
				go lconn.Start()
				atomic.AddUint64(&connId, 1)
			}
		}()

		select {
		case <-s.exitCh:
			if err := listen.Close(); err != nil {
				log.Println("listener close err: %v", err)
			}
		}
	}()
}

func (s *Server) Stop() {
	log.Println("[STOP] Lion server , name %s", s.Name)
	s.exitCh <- struct{}{}
	close(s.exitCh)
}

func (s *Server) apply(opts ...Option) {
	for _, opt := range opts {
		opt(s)
	}
}
