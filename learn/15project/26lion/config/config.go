package config

import (
	"encoding/json"
	"github.com/spf13/viper"
	"log"
	"os"
	"sync"
)

const (
	LOCAL = "local"
	DEV   = "edv"
	LINE  = "line"
)

var (
	pv     *proxyViper
	pvOnce sync.Once
)

type proxyViper struct {
	v   *viper.Viper
	err error
}

func newProxyViper() *proxyViper {
	return &proxyViper{}
}

func InitializeProxyViper() error {
	pvOnce.Do(func() {
		pv = newProxyViper()
		filePath, err := getPath()
		if err != nil {
			pv.err = err
			return
		}
		newViper := viper.New()
		newViper.SetConfigFile(filePath)
		if err = newViper.ReadInConfig(); err != nil {
			pv.err = err
			return
		}
		pv.v = newViper
		log.Println("config init ok..............................")
	})
	return pv.err
}

func Generate(key string, dst interface{}) error {

	if err := InitializeProxyViper(); err != nil {
		return err
	}
	ret := pv.v.Get(key)
	s, err := json.Marshal(ret)
	if err != nil {
		return err
	}
	return json.Unmarshal(s, &dst)
}

func getPath() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	p := path + "/conf/"
	switch getEnv() {
	case LOCAL:
		return p + LOCAL + ".yaml", nil
	case DEV:
		return p + DEV + ".yaml", nil
	case LINE:
		return p + LINE + ".yaml", nil
	default:
		return p + LOCAL + ".yaml", nil
	}
}

func getEnv() string {
	return os.Getenv("c_env")
}
