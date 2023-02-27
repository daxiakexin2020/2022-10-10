package server

import (
	"fmt"
	"os"
	"testing"
)

func Test01(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	Start(path)
}
