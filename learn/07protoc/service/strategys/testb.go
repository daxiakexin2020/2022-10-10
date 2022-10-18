package strategys

import "fmt"

type TestB struct {
	General
}

func NewTestB() *TestB {
	return &TestB{}
}

func (tb *TestB) Decode(filepath string) error {
	fmt.Println("策略B实现了 decode接口")
	return nil
}
