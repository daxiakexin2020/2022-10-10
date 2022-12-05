package register_center

import (
	"14gateway/register_center/etcd/service"
	"context"
	"errors"
)

type regTyep string

const (
	ETCD   regTyep = "etcd"
	CONSOL regTyep = "consol"
)

type Register interface {
	Name() string
	Get(ctx context.Context, key string) ([]string, error)
	Put(ctx context.Context, key string, val string, options ...interface{}) error
	Delete(ctx context.Context, key string) error
}

type RegisteServer struct {
	RType regTyep  `json:"rtype"`
	reg   Register `json:"reg"`
}

func NewRegisteServer(rtyep regTyep) (*RegisteServer, error) {
	rs := &RegisteServer{RType: rtyep}
	switch rtyep {
	case ETCD:
		rs.reg = service.GetEtcdClient()
		return rs, nil
	case CONSOL:
		return nil, errors.New("CONSOL注册中心正在建设中～～～")
	default:
		return nil, errors.New("没有注册中心可以使用")
	}
}

func (rs *RegisteServer) Put(key string, val string, options ...interface{}) error {
	return rs.reg.Put(context.Background(), key, val, options...)
}

func (rs *RegisteServer) Get(key string) ([]string, error) {
	return rs.reg.Get(context.Background(), key)
}

func (rs *RegisteServer) Delete(key string) error {
	return rs.reg.Delete(context.Background(), key)
}
