package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {

	//返回2个字节切片的字典序的大小， a==b>0   a<b >-1  a>b >+1
	compare := bytes.Compare([]byte("a"), []byte("b"))
	fmt.Println("compare:", compare)

	//判断2个字节数组，是否完全相同
	equal := bytes.Equal([]byte("AB"), []byte("AB"))
	fmt.Println("equal:", equal)

	//判断是否包含前缀字节，从第一个开始判断     HasSuffix 后缀
	prefix := bytes.HasPrefix([]byte("abc"), []byte("ab"))
	fmt.Println("hasPrefix:", prefix)

	//判断b切片是否包含a切片
	contains := bytes.Contains([]byte("abc"), []byte("a"))
	fmt.Println("contains:", contains)

	//返回将所有字母转为小写的切片   ToUpper 大写
	lower := bytes.ToLower([]byte("aBcD"))
	fmt.Println("toLower:", string(lower))

	//返回b个a形式的串联的切片
	repeat := bytes.Repeat([]byte("ab"), 3)
	fmt.Println("repeat:", repeat, string(repeat)) //[97 98 97 98 97 98] ababab

	//将byte中前后的包含的字符串去掉，仅限于前后的，不是所有的
	trim := bytes.Trim([]byte("abcdabcda"), "a")
	fmt.Println("trim:", trim, string(trim)) //[98 99 100 97 98 99 100] bcdabcd

	//返回将s前后端所有满足f的unicode码值都去掉的子切片。（共用底层数组）
	trimfuncByte := []byte("abcdabcd")
	trimFunc := bytes.TrimFunc(trimfuncByte, func(r rune) bool {
		fmt.Println("r:", string(r))
		if r == 'a' {
			return false
		}
		return true
	})
	fmt.Println("trimFunc:", trimFunc, string(trimFunc))

	//NewReader创建一个从b中读取数据的Reader。
	b := []byte{'a', 'b', 'c'}
	reader := bytes.NewReader(b)
	fmt.Println("newReader:", reader)

	//从b中读取字节，放入r中
	r := make([]byte, 3)
	n, err := reader.Read(r)
	fmt.Println("reader read:", n, err, r, string(r))

	//NewBuffer使用buf作为初始内容创建并初始化一个Buffer。
	//本函数用于创建一个用于读取已存在数据的buffer；也用于指定用于写入的内部缓冲的大小，此时，buf应为一个具有指定容量但长度为0的切片。buf会被作为返回值的底层缓冲切片。
	buf := make([]byte, 0, 10)
	fmt.Println("len , cap", len(buf), cap(buf))
	buffer := bytes.NewBuffer(buf)

	var _ io.Reader = (buffer)

	fmt.Println("newBuffer:", buffer)

	//write string
	buffer.WriteString("abc")
	buffer.WriteString("def")
	s := buffer.String()
	fmt.Println("reader buffer write string:", s)

	buffer.WriteString("ghi")
	s2 := buffer.String()
	fmt.Println("reader buffer write string2:", s2)

	//清空底层buf，否则会一直加入[]byte
	buffer.Reset()
	s3 := buffer.String()
	fmt.Println("reader buffer write string3:", s3)

}
