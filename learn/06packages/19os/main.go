package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {

	//返回内核提供的主机名称
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("hostname err:", err)
	}
	fmt.Println("hostname:", name)

	//Getenv检索并返回名为key的环境变量的值。如果不存在该环境变量会返回空字符串。
	getenv := os.Getenv("PATH")
	fmt.Println("get env:", getenv)

	//返回调用者的uid
	getuid := os.Getuid()
	fmt.Println("getuid:", getuid)

	//Getwd返回一个对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd会返回其中一个。
	dir, _ := os.Getwd()
	fmt.Println("getwd:", dir)

	//Mkdir使用指定的权限和名称创建一个目录。如果出错，会返回*PathError底层类型的错误。会默认使用当前项目路径
	err = os.Mkdir("b", 0644)
	fmt.Println("mkdir err:", err)

	//NewFile函数是os包用于新建文件的函数。NewFile并不是真正创建了一个文件，而是新建了文件但并不保存，返回新建后文件的指针。
	file := os.NewFile(0, "a")
	n, err := file.WriteString("aaaaaa")
	fmt.Println("newfile:", n, err)

	//读取文件内容，从指定位置读取
	//ReadAt从指定的位置（相对于文件开始位置）读取len(b)字节数据并写入b。它返回读取的字节数和可能遇到的任何错误。当n<len(b)时，本方法总是会返回错误；todo 如果是因为到达文件结尾，返回值err会是io.EOF。
	path := dir + "/learn/06packages/19os/test.txt"
	open, err := os.Open(path)
	if err != nil {
		fmt.Println("open file err:", err)
	} else {
		defer open.Close()
		var bs []byte
		var offest int64
		for {
			b := make([]byte, 1)
			at, err := open.ReadAt(b, offest)
			fmt.Println("readat:", at, string(b), err)
			if err == io.EOF {
				break
			}
			offest += int64(at)
			bs = append(bs, b...)
		}
		fmt.Println("readat over:", string(bs))
	}

	//执行外部命令
	command := exec.Command("ls")
	err = command.Run()
	fmt.Println("exec run err : ", err)

	//执行命令并返回标准输出的切片。
	output, err := command.Output()
	if err != nil {
		fmt.Println("exec output err:", err)
	} else {
		fmt.Println("output:", string(output))
	}

	//监听终止信号,一般用于优雅关闭服务¬
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, os.Kill)
L:
	for {
		select {
		case <-c:
			fmt.Println("终止了.....")
			break L
		}
	}
	fmt.Println("over")
}
