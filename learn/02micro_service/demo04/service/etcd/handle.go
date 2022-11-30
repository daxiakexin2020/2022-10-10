package etcd

import (
	etcd_service "03middleware/etcd/service"
	"context"
	"errors"
	"time"
)

type Etcd struct {
	Address []string `json:"address"`
	Proxy   *etcd_service.Etcd
}

func NewEtcd(addr []string) (*Etcd, error) {
	es, err := etcd_service.NewEtcd(addr, etcd_service.WithDialTimeout(1*time.Second))
	if err != nil {
		return nil, errors.New("连接etcd服务失败")
	}
	return &Etcd{Address: addr, Proxy: es}, nil
}

func (e *Etcd) Name() string {
	return "etcd"
}

func (e *Etcd) Get(ctx context.Context, key string) ([]string, error) {
	return e.Proxy.Get(ctx, key)
}

func (e *Etcd) Put(ctx context.Context, key string, val string) error {
	return e.Proxy.Put(ctx, key, val)
}
