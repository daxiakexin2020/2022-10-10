package demo

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestDemo(t *testing.T) {

	dir, err := os.Getwd()
	if err != nil {
		log.Panic("指定模版文件不存在")
	}
	path := dir + "/template/test.go.tmp"
	file, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Panic("file err", err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	fmt.Println(string(content))
	fmt.Println("b", content)
}
