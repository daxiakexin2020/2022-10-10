package config

import (
	"log"
	"sync"
	"time"
)

var (
	jwtConfig JwtConfig
	once      sync.Once
)

func init() {
	//makeJwtConfig()
}

type JwtConfig struct {
	TokenSecret  string        `json:"token_secret"`
	TokenIssuer  string        `json:"token_issuer"`
	TokenTimeout time.Duration `json:"token_timeout"`
}

func (jc *JwtConfig) CName() string {
	return "jwt"
}

func makeJwtConfig() JwtConfig {
	once.Do(func() {
		jc := JwtConfig{}
		if err := Generate(jc.CName(), &jc); err != nil {
			log.Fatalf("读取%s配置出错%v", jc.CName(), err)
		}
		jwtConfig = jc
	})
	return jwtConfig
}

func GetJwtConfig() JwtConfig {
	return makeJwtConfig()
}
