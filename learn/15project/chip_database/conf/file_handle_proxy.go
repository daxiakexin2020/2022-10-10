package conf

import (
	"log"
	"sync"
)

var (
	fileHandleProxyConfig *FileHandleProxyConfig
	fileHandleProxyOnce   sync.Once
)

type FileHandleProxyConfig struct {
	Base string `json:"base"`
}

func (c *FileHandleProxyConfig) CName() string {
	return "file_handle_proxy"
}

func makeFileHandleProxyConfig() {
	fileHandleProxyOnce.Do(func() {
		c := &FileHandleProxyConfig{}
		if err := generate(c.CName(), c); err != nil {
			log.Fatalf("read %s config err%v\n", c.CName(), err)
		}
		fileHandleProxyConfig = c
		log.Printf("read %s config ok::::::::::::::::::::::::%+v\n", c.CName(), c)
	})
}

func GetFileHandleProxyOnce() *FileHandleProxyConfig {
	if fileHandleProxyConfig == nil {
		makeFileHandleProxyConfig()
	}
	return fileHandleProxyConfig
}
