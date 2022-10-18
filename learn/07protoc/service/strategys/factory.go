package strategys

import (
	"07protoc/service"
	"errors"
	"sync"
)

var OFactory *Factory

type Factory struct {
	IsRegister bool
	list       map[string]service.Parser
}

func RegisterStrategy() {
	once := sync.Once{}
	once.Do(func() {
		if OFactory == nil || !OFactory.IsRegister {
			list := make(map[string]service.Parser)
			list["testa"] = NewTestA()
			list["testb"] = NewTestB()
			OFactory = &Factory{
				IsRegister: true,
				list:       list,
			}
		}
	})
}

func (f *Factory) MakeStrategy(name string) (service.Parser, error) {
	if paser, ok := OFactory.list[name]; ok {
		return paser, nil
	}
	return nil, errors.New("暂不支持此种策略")
}
