package v1

import (
	"os"
	"testing"
)

func TestInitializeBigFile(t *testing.T) {
	dir, _ := os.Getwd()
	s := NewServer(dir + "/test_files")
	s.InitializeBigFile()
}

func TestSpiltSmallFiles(t *testing.T) {
	dir, _ := os.Getwd()
	s := NewServer(dir + "/test_files")
	s.SpiltSmallFiles()
}

func TestEverySmallFileCount(t *testing.T) {
	dir, _ := os.Getwd()
	s := NewServer(dir + "/test_files")
	s.EverySmallFileCount()
}
