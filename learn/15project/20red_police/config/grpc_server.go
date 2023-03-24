package config

import (
	"log"
	"sync"
)

var (
	grpcServerConfig GrpcServerConfig
	grpcServerOnce   sync.Once
)

type GrpcServerConfig struct {
	Addr         string `json:"addr"`
	Port         int    `json:"port"`
	ReadDeadLine int    `json:"read_dead_line"`
}

func (gs *GrpcServerConfig) CName() string {
	return "grpc_server"
}

func init() {
	//makeGrpcServerConfig()
}

func makeGrpcServerConfig() GrpcServerConfig {
	grpcServerOnce.Do(func() {
		gsc := GrpcServerConfig{}
		if err := Generate(gsc.CName(), &gsc); err != nil {
			log.Fatalf("读取%s配置出错%v", gsc.CName(), err)
		}
		grpcServerConfig = gsc
	})
	return grpcServerConfig
}

func GetGrpcServerConfig() GrpcServerConfig {
	return makeGrpcServerConfig()
}
