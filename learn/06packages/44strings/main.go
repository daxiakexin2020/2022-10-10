package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcdefabaa"

	//子串sep在字符串s中第一次出现的位置，不存在则返回-1。LastIndex
	index := strings.Index(str, "a")
	fmt.Printf("Index()=%d\n", index)

	//将后端所有的包含的字符的去掉 ，从右侧看
	right := strings.TrimRight(str, "a")
	fmt.Printf("TrimRight=%s\n", right)

	//返回将s前后端所有cutset包含的utf-8码值都去掉的字符串。
	trim := strings.Trim(str, "a")
	fmt.Printf("Trim=%s\n", trim)

	//返回将字符串按照空白（unicode.IsSpace确定，可以是一到多个连续的空白字符）分割的多个字符串。如果字符串全部是空白或者是空字符串的话，会返回空切片。
	str2 := "a	b		c    d	"
	fields := strings.Fields(str2)
	fmt.Printf("Fields=%v,Len=%d\n", fields, len(fields))

	//注意与Fields()的区别
	split := strings.Split(str2, "	")
	fmt.Printf("Split=%v,Len=%d\n", split, len(split))
	fmt.Println(fields[3], "===", split[3])
}
