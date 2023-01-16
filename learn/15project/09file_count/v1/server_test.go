package v1

import (
	"os"
	"testing"
)

func Test01(t *testing.T) {
	dir, _ := os.Getwd()
	s := NewServer(dir + "/test_files")
	s.InitializeBigFile()
}
