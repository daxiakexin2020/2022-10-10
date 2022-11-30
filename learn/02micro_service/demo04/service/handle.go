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

type Service struct {
	ServerName string   `json:"server_name"`
	Address    []string `json:"address"`
	RType      regTyep
	reg        Register
}

func NewService(serviceName string, addreess []string, rtype regTyep) (*Service, error) {
	service := &Service{
		ServerName: serviceName,
		Address:    addreess,
		RType:      rtype,
	}

	switch rtype {
	case ETCD:
		r, err := etcd.NewEtcd([]string{"127.0.0.1:2379"})
		if err != nil {
			return nil, err
		}
		service.reg = r
	default:
		return nil, errors.New("没有注册中心可用")
	}

	return service, nil
}

func (s *Service) Register() error {
	key, val := s.Gerenral()
	return s.reg.Put(context.TODO(), key, val)
}

func (s *Service) Get(key string) ([]string, error) {
	return s.reg.Get(context.TODO(), key)
}

func (s *Service) Gerenral() (string, string) {
	addrs := strings.Join(s.Address, "-")
	key := fmt.Sprintf("%s/%s", s.ServerName, addrs)
	val := fmt.Sprintf("%s", addrs)
	return key, val
}
