package conf

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type CRedis struct {
	Redis Option
}

type Option struct {
	Host     string `json:"host"`
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

var v *viper.Viper

func Handle() error {
	initConfig()
	return nil
}

func initConfig() {
	root, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	v = viper.New()
	v.AddConfigPath(root + "redis/config")
	v.SetConfigType("yaml")
	v.SetConfigName("test")
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
