package main

import (
	"flag"
	"fmt"
)

func main() {

	//将输入的变量保存在一个int变量中
	var intflag int
	flag.IntVar(&intflag, "iflag", 1, "inout int flag")
	flag.Parse()
	fmt.Println(flag.Args())
	fmt.Println("iflag:", intflag)
}
