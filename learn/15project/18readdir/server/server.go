package server

import (
	"fmt"
	"io/ioutil"
	"os"
)

/**
读取指定文件夹下，所有文件的个数，以及大小
*/

func Start(dir string) {
	readAllDirs(dir)
}

func readAllDirs(dir string) []os.FileInfo {
	readDir, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return readDir
}
