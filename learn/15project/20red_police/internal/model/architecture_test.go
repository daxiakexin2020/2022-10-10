package model

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	initArchitecture := InitArchitecture()
	strings := initArchitecture.List()
	for _, name := range strings {
		fmt.Println("name:", name)
		architectureArm, err := initArchitecture.FetchArchitectureArm(name)
		fmt.Println("arm list", architectureArm, err)
	}
}
