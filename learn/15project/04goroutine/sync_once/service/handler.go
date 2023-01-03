package service

import (
	"fmt"
	"sync"
)

type Config struct {
	Server string
	Port   int32
}

var (
	conce  sync.Once
	config *Config
)

func doOnce(wg *sync.WaitGroup) {
	defer wg.Done()
	conce.Do(func() {
		config = &Config{
			Server: "locahost",
			Port:   8899,
		}
		fmt.Println("config初始化ok")
	})
}

func do(wg *sync.WaitGroup) {
	defer wg.Done()
	config = &Config{
		Server: "locahost",
		Port:   8899,
	}
	fmt.Println("config初始化ok")
}
