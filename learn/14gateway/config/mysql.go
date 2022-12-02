package config

import (
	"log"
	"sync"
)

var (
	mysqlConfig *MysqlConfig
	mysqlOnce   sync.Once
)

type MysqlConfig struct {
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (mc *MysqlConfig) CName() string {
	return "mysql"
}

func makeMysqlConfig() {
	mysqlOnce.Do(func() {
		mc := &MysqlConfig{}
		if err := Generate(mc.CName(), mc); err != nil {
			log.Fatalf("读取%s配置出错%v", mc.CName(), err)
		}
		mysqlConfig = mc
	})
}

func GetMysqlConfig() *MysqlConfig {
	if mysqlConfig == nil {
		makeMysqlConfig()
	}
	return mysqlConfig
}
