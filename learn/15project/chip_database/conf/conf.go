package conf

import (
	"encoding/json"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
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

func initializeProxyViper() error {
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
	})
	return pv.err
}

func generate(key string, dst interface{}) error {
	if err := initializeProxyViper(); err != nil {
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

	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	abs, err := filepath.Abs(currentDir)
	if err != nil {
		return "", err
	}
	path := filepath.Join(abs, "../")
	p := path + "/config/"
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
