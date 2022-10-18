package strategys

import "fmt"

type TestA struct {
	General
}

func (ta *TestA) Decode(filepath string) error {
	fmt.Println("策略A实现了decode接口")
	return nil
}

func NewTestA() *TestA {
	return &TestA{}
}
