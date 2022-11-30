package service

import (
	"context"
	"demo04/service/etcd"
	"errors"
	"fmt"
	"strings"
)

type regTyep string

const (
	ETCD   regTyep = "etcd"
	CONSOL regTyep = "consol"
)

type Register interface {
	Name() string
	Get(ctx context.Context, key string) ([]string, error)
	Put(ctx context.Context, key string, val string) error
}

type RegisteServer struct {
	RType regTyep  `json:"rtype"`
	reg   Register `json:"reg"`
}

func NewRegisteServer(addr []string, rtyep regTyep) (*RegisteServer, error) {
	rs := &RegisteServer{RType: rtyep}
	switch rtyep {
	case ETCD:
		r, err := etcd.NewEtcd(addr)
		if err != nil {
			return nil, err
		}
		rs.reg = r
		return rs, nil
	default:
		return nil, errors.New("没有注册中心可以使用")
	}
}

func (rs *RegisteServer) Register(serverName string, addr []string) error {
	return rs.reg.Put(context.Background(), rs.makeKey(serverName), rs.makeValue(addr))
}

func (rs *RegisteServer) Get(key string) ([]string, error) {
	return rs.reg.Get(context.TODO(), rs.makeKey(key))
}

func (rs *RegisteServer) makeKey(key string) string {
	return fmt.Sprintf("service_%s", key)
}

func (rs *RegisteServer) makeValue(val []string) string {
	return strings.Join(val, "-")
}
