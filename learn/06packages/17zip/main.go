package main

import (
	"archive/zip"
	"os"
)

func main() {

	//创建一个zip文件
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	name := path + "/" + "learn/06packages/17zip/file/test.zip"
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	//这里接收io.Writer,包装成zip的Writer
	writer := zip.NewWriter(file)

	//创造出io.Writer
	zfile, err := writer.Create(name)
	zfile.Write([]byte("abc"))
	defer writer.Close()
}
