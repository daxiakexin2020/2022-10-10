package service

import (
	"go.etcd.io/etcd/client/v3"
	"time"
)

type Etcd struct {
	Client    *clientv3.Client
	Config    clientv3.Config
	Endpoints []string
}

type Option func(e *Etcd)

func (e *Etcd) apply(opts []Option) {
	for _, opt := range opts {
		opt(e)
	}
}

func WithDialTimeout(dialTimeout time.Duration) Option {
	return func(e *Etcd) {
		e.Config.DialTimeout = dialTimeout
	}
}

func NewEtcd(endpoints []string, opts ...Option) (*Etcd, error) {
	etcd := &Etcd{Config: clientv3.Config{}}
	etcd.apply(opts)
	etcd.Config.Endpoints = endpoints
	cli, err := clientv3.New(etcd.Config)
	if err != nil {
		return nil, err
	}
	etcd.Client = cli
	return etcd, nil
}
