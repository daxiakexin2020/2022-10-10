package config

import (
	"log"
	"sync"
)

var (
	sverConfig ServerConfig
	serverOnce sync.Once
)

type ServerConfig struct {
	Name      string `json:"name"`
	IPVersion string `json:"ip_version"`
	Ip        string `json:"ip"`
	Port      int    `json:"port"`
}

func (gs *ServerConfig) CName() string {
	return "server"
}

func makeServerConfig() ServerConfig {
	serverOnce.Do(func() {
		gsc := ServerConfig{}
		if err := Generate(gsc.CName(), &gsc); err != nil {
			log.Fatalf("读取%s配置出错%v", gsc.CName(), err)
		}
		sverConfig = gsc
	})
	return sverConfig
}

func GetServerConfig() ServerConfig {
	return makeServerConfig()
}
