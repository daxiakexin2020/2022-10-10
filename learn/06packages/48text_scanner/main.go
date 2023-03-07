package main

import (
	"fmt"
	"io"
	"os"
	"text/scanner"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := dir + "/learn/06packages/48text_scanner/test.txt"
	file, err := os.OpenFile(path, os.O_RDWR, 0777)
	if err != nil {
		panic(err)
	}
	var _ io.Reader = (file)
	defer file.Close()

	var s scanner.Scanner
	inits := s.Init(file)

	for {
		next := inits.Scan()
		if next == scanner.EOF {
			return
		}
		text := inits.TokenText()
		fmt.Printf("read s=%s\n", text)
	}
}
