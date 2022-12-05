package service

import (
	"14gateway/config"
	"fmt"
)

func InitEtcd() error {
	ec := config.GetEtcdConfig()
	options := make([]Option, 0)
	if ec.Username != "" {
		options = append(options, WithUsername(ec.Username))
	}
	if ec.Password != "" {
		options = append(options, WithPassword(ec.Password))
	}
	if ec.DialTimeout > 0 {
		options = append(options, WithDialTimeout(ec.DialTimeout))
	}
	if ec.DialKeepAliveTime > 0 {
		options = append(options, WithDialKeepAliveTime(ec.DialKeepAliveTime))
	}
	_, err := NewEtcd(ec.Endpoints, options...)

	if err != nil {
		return fmt.Errorf("连接etcd失败%v", err)
	}
	return nil
}
