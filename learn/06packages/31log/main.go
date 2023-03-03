package main

import (
	"io"
	"log"
	"os"
)

func main() {
	log.Println("log输出到标准输出....")

	//自定义输出io.Writer 输出到文件
	pwd, _ := os.Getwd()
	file, err := os.OpenFile(pwd+"/learn/06packages/31log/a.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panicln("file open err:", err)
	}
	var _ io.Writer = (file)
	l := log.New(file, "[test]", 5)
	l.Println("写入文件")
}
