package config

import (
	"log"
	"sync"
)

var (
	msgHandleConfig MsgHandleConfig
	msgHandleOnce   sync.Once
)

type MsgHandleConfig struct {
	WorkPoolSize uint32 `json:"work_pool_size"`
}

func (gs *MsgHandleConfig) CName() string {
	return "msghandle"
}

func makeMsgHandleConfig() MsgHandleConfig {
	serverOnce.Do(func() {
		c := MsgHandleConfig{}
		if err := Generate(c.CName(), &c); err != nil {
			log.Fatalf("读取%s配置出错%v", c.CName(), err)
		}
		msgHandleConfig = c
	})
	return msgHandleConfig
}

func GetMsgHandleConfig() MsgHandleConfig {
	return makeMsgHandleConfig()
}
