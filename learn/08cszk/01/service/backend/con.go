package backend

import "fmt"

type Config struct {
	Addr string
	Port int64
}

func NewCon() *Config {
	return &Config{}
}

func (c *Config) Name() string {
	return "conh"
}

func (c *Config) InitService() error {
	c.Addr = "127.0.0.1"
	c.Port = 80
	fmt.Println("*******************CONH服务	init	ok******************")
	return nil
}
