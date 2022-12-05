package service

import (
	"go.etcd.io/etcd/client/v3"
	"sync"
	"time"
)

var (
	etcdClient *Etcd
	eoncy      sync.Once
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

func WithUsername(username string) Option {
	return func(e *Etcd) {
		e.Config.Username = username
	}
}

func WithPassword(password string) Option {
	return func(e *Etcd) {
		e.Config.Password = password
	}
}

func WithDialTimeout(dialTimeout int64) Option {
	return func(e *Etcd) {
		e.Config.DialTimeout = time.Second * time.Duration(dialTimeout)
	}
}

func WithDialKeepAliveTime(dialKeepAliveTime int64) Option {
	return func(e *Etcd) {
		e.Config.DialKeepAliveTime = time.Second * time.Duration(dialKeepAliveTime)
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
	etcdClient = etcd
	return etcdClient, nil
}

func GetEtcdClient() *Etcd {
	return etcdClient
}
