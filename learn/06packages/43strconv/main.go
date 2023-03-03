package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "F"

	//返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	parseBool, err := strconv.ParseBool(str)
	fmt.Printf("ParseBool=%t,err=%v\n", parseBool, err)

	//返回字符串表示的整数值，接受正负号。
	//base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
	//bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；
	//返回的err是*NumErr类型的，如果语法有误，err.Error = ErrSyntax；如果结果超出类型范围err.Error = ErrRange。
	str2 := "1"
	i, err := strconv.ParseInt(str2, 0, 0)
	fmt.Printf("ParseInt=%d,err=%v\n", i, err)

	//Atoi是ParseInt(s, 10, 0)的简写。
	atoi, err := strconv.Atoi(str2)
	fmt.Printf("Atoi=%d,err=%v\n", atoi, err)

	//返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
	idata := int64(611)
	formatInt := strconv.FormatInt(idata, 10)
	fmt.Printf("Format=%s\n", formatInt)
}
