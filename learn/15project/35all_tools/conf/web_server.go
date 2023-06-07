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

func (ec *WebServerConfig) CName() string {
	return "web_server"
}

func makeWebServerConfig() {
	webServerOnce.Do(func() {
		wsc := &WebServerConfig{}
		if err := Generate(wsc.CName(), wsc); err != nil {
			log.Fatalf("读取%s配置出错%v\n", wsc.CName(), err)
		}
		webServerConfig = wsc
		log.Printf("读取%s配置::::::::::::::::::::::::%v\n", wsc.CName(), wsc)
	})
}

func GetWebServerConfig() *WebServerConfig {
	if webServerConfig == nil {
		makeWebServerConfig()
	}
	return webServerConfig
}
