package conf

import (
	"errors"
	"github.com/spf13/viper"
	"log"
	"os"
)

type MysqlConfig struct {
	Mysql config
}

type config struct {
	Host     string `json:"host" mapstructure:host`
	Port     string `json:"port" mapstructure:port`
	Database string `json:"database" mapstructure:database`
	Username string `json:"username" mapstructure:username`
	Password string `json:"password" mapstructure:password`
}

var v *viper.Viper
var mc *MysqlConfig

func Handle() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}
	mc = &MysqlConfig{}
	v.Unmarshal(mc)
}

func initConfig() error {

	v = viper.New()
	v.SetConfigName("test")
	v.SetConfigType("yaml")

	root, err := os.Getwd()
	if err != nil {
		return errors.New("root path error")
	}

	v.AddConfigPath(root + "/config")

	if err = v.ReadInConfig(); err != nil {
		log.Fatal("Read config error", err)
	}

	return nil
}

func GetMysqlConfig() (*MysqlConfig, error) {
	if mc == nil {
		return nil, errors.New("未初始化配置")
	}
	return mc, nil
}
