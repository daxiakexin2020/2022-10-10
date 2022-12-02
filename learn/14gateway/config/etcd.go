package config

import (
	"log"
	"sync"
	"time"
)

var (
	etcdConfig *EtcdConfig
	etcdOnce   sync.Once
)

type EtcdConfig struct {
	Endpoints         []string      `json:"endpoints" mapstructure:"endpoints"`
	Username          string        `json:"username" mapstructure:"username"`
	Password          string        `json:"password" mapstructure:"password"`
	DialTimeout       time.Duration `json:"dial_timeout" mapstructure:"dialTimeout"`
	DialKeepAliveTime time.Duration `json:"dial_keep_alive_time" mapstructure:"dialKeepAliveTime"`
}

func (ec *EtcdConfig) CName() string {
	return "etcd"
}

func makeEtcdConfig() {
	etcdOnce.Do(func() {
		ec := &EtcdConfig{}
		if err := Generate(ec.CName(), ec); err != nil {
			log.Fatalf("读取%s配置出错%v", ec.CName(), err)
		}
		etcdConfig = ec
	})
}

func GetEtcdConfig() *EtcdConfig {
	if etcdConfig == nil {
		makeEtcdConfig()
	}
	return etcdConfig
}
