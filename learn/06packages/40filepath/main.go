package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, _ := os.Getwd()
	root := dir + "/learn/06packages/40filepath"
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		fmt.Println("path:", path)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
