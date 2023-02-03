package generate

import "testing"

type TestStruct struct {
	Name string `zorm:ame`
}

func TestB(t *testing.T) {
	data := TestStruct{
		Name: "zz",
	}
	B(data)
}
