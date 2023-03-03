package main

import (
	"fmt"
	"path"
)

func main() {
	p := "/User/zz/test.txt"

	//判断一个路径是否是一个绝对路径
	abs := path.IsAbs(p)
	fmt.Printf("abs=%t\n", abs)

	//Split函数将路径从最后一个斜杠后面位置分隔为两个部分（dir和file）并返回。如果路径中没有斜杠，函数返回值dir会设为空字符串，file会设为path。两个返回值满足path == dir+file。
	dir, file := path.Split(p)
	fmt.Printf("dir=%s,file=%s\n", dir, file)

	//Join函数可以将任意数量的路径元素放入一个单一路径里，会根据需要添加斜杠。结果是经过简化的，所有的空字符串元素会被忽略。
	//相当于把给的各个分路径，组合在一起，按照传入的顺序
	p1 := "/User"
	p2 := "demo.txt"
	join := path.Join(p1, p2)
	fmt.Printf("join=%s\n", join) //join=/User/demo.txt

	//返回路径部分，会截取最后一个/之前的路径
	//Dir返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。
	//在使用Split去掉最后一个元素后，会简化路径并去掉末尾的斜杠。如果路径是空字符串，会返回"."；如果路径由1到多个斜杠后跟0到多个非斜杠字符组成，会返回"/"；其他任何情况下都不会返回以斜杠结尾的路径。
	s := path.Dir(p)
	fmt.Printf("dir=%s\n", s)

	//会截取路径中最后一个元素，不管是文件还是路径
	//Base函数返回路径的最后一个元素。在提取元素前会求掉末尾的斜杠。如果路径是""，会返回"."；如果路径是只有一个斜杆构成，会返回"/"。
	base := path.Base(p)
	fmt.Printf("base=%s", base)
}
