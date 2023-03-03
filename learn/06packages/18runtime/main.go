package main

import (
	"fmt"
	"runtime"
)

func main() {

	//GOROOT返回Go的根目录。如果存在GOROOT环境变量，返回该变量的值；否则，返回创建Go时的根目录。
	goroot := runtime.GOROOT()
	fmt.Println("go root : ", goroot)

	//GOMAXPROCS设置可同时执行的最大CPU数，并返回先前的设置。 若 n < 1，它就不会更改当前设置。本地机器的逻辑CPU数可通过 NumCPU 查询。本函数在调度程序优化后会去掉。
	gomaxprocs := runtime.GOMAXPROCS(3)
	fmt.Println("gomaxprocs : ", gomaxprocs)

	//执行一次垃圾回收
	runtime.GC()

	a := make(map[string]string)
	a["1"] = "1"
	fmt.Println(a)

	record := runtime.MemProfileRecord{}
	//正在使用的字节数
	bytes := record.InUseBytes()
	fmt.Println("in use bytes : ", bytes)

	//Stack 返回格式化的go程的调用栈踪迹。 对于每一个调用栈，它包括原文件的行信息和PC值；对go函数还会尝试获取调用该函数的函数或方法，及调用所在行的文本。
	rbyte := make([]byte, 1024)
	stack := runtime.Stack(rbyte, true)
	fmt.Println("stack:", stack, string(rbyte))
}
