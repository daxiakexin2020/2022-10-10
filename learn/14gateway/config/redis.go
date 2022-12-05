package config

import (
	"log"
	"sync"
)

var (
	redisConfig *RedisConfig
	redisOnce   sync.Once
)

type RedisConfig struct {
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
	DB       int    `json:"db"`
	Password string `json:"password"`
	Prefix   string `json:"prefix"`
	PoolSize int    `json:"pool_size"`
}

func (rc *RedisConfig) CName() string {
	return "redis"
}

func makeRedisConfig() {
	redisOnce.Do(func() {
		rc := &RedisConfig{}
		if err := Generate(rc.CName(), rc); err != nil {
			log.Fatalf("读取%s配置出错%v", rc.CName(), err)
		}
		redisConfig = rc
	})
}

func GetRedisConfig() *RedisConfig {
	if redisConfig == nil {
		makeRedisConfig()
	}
	return redisConfig
}
