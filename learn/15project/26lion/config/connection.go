package config

import (
	"log"
	"sync"
)

var (
	connctionConfig ConnctionConfig
	connctioOnce    sync.Once
)

type ConnctionConfig struct {
	IoReaderBufferSize uint32 `json:"io_reader_buffer_size"`
}

func (cc *ConnctionConfig) CName() string {
	return "connection"
}

func makeConnctionConfig() ConnctionConfig {
	connctioOnce.Do(func() {
		c := ConnctionConfig{}
		if err := Generate(c.CName(), &c); err != nil {
			log.Fatalf("读取%s配置出错%v", c.CName(), err)
		}
		connctionConfig = c
	})
	return connctionConfig
}

func GetConnctionConfigConfig() ConnctionConfig {
	return makeConnctionConfig()
}
