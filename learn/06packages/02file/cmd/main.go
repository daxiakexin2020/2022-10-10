package main

import (
	"02file/service"
	"fmt"
)

func main() {
	testHandle()
}

func testHandle() {
	s := service.FilePathJoin()
	s = service.TmpTest()
	fmt.Printf("s=%s", s)
}
