package config

import (
	"log"
	"sync"
	"time"
)

var (
	jwtConfig *JwtConfig
	once      sync.Once
)

type JwtConfig struct {
	TokenSecret  string        `json:"token_secret"`
	TokenIssuer  string        `json:"token_issuer"`
	TokenTimeout time.Duration `json:"token_timeout"`
}

func (jc *JwtConfig) CName() string {
	return "jwt"
}

func makeJwtConfig() {
	once.Do(func() {
		jc := &JwtConfig{}
		if err := Generate(jc.CName(), jc); err != nil {
			log.Fatalf("读取%s配置出错%v", jc.CName(), err)
		}
		jwtConfig = jc
	})
}

func GetJwtConfig() *JwtConfig {
	if jwtConfig == nil {
		makeJwtConfig()
	}
	return jwtConfig
}
