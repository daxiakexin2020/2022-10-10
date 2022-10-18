package strategys

import (
	"07protoc/service"
	"fmt"
)

type General struct {
	service.Parser
}

func NewGeneral(parser service.Parser) *General {
	return &General{}
}

func (g *General) Decode(filepath string) error {
	fmt.Println("general 实现了decode接口")
	return nil
}
