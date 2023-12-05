package conf

import (
	"log"
	"sync"
)

var (
	webServerConfig *WebServerConfig
	webServerOnce   sync.Once
)

type WebServerConfig struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (c *WebServerConfig) CName() string {
	return "web_server"
}

func makeWebServerConfig() {
	webServerOnce.Do(func() {
		c := &WebServerConfig{}
		if err := generate(c.CName(), c); err != nil {
			log.Fatalf("read %s config err%v\n", c.CName(), err)
		}
		webServerConfig = c
		log.Printf("read %s config ok::::::::::::::::::::::::%+v\n", c.CName(), c)
	})
}

func GetWebServerConfig() *WebServerConfig {
	if webServerConfig == nil {
		makeWebServerConfig()
	}
	return webServerConfig
}
