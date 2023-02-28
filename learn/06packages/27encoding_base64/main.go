package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	//使用给出的字符集生成一个*Encoding，字符集必须是64字节的字符串。
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	encoding := base64.NewEncoding(str)
	fmt.Println("new encoding:", encoding)

	//将src的数据编码后存入dst，最多写EncodedLen(len(src))字节数据到dst，并返回写入的字节数。
	// 函数会把输出设置为4的倍数，因此不建议对大数据流的独立数据块执行此方法，使用NewEncoder()代替。
	src := []byte("abc")
	dst := make([]byte, encoding.EncodedLen(len(src)))
	encoding.Encode(dst, src)
	fmt.Println("encoding encode:", dst, string(dst))

	//返回直接将src进行编码后的字符串
	toString := encoding.EncodeToString(src)
	fmt.Println("toString:", toString)

	//将src的数据解码后存入dst，最多写DecodedLen(len(src))字节数据到dst，并返回写入的字节数。
	//如果src包含非法字符，将返回成功写入的字符数和CorruptInputError。换行符（\r、\n）会被忽略。
	dst2 := make([]byte, encoding.DecodedLen(len(dst)))
	n, err := encoding.Decode(dst2, dst)
	fmt.Println("encoding decode:", n, err, dst2, string(dst2))

	//使用默认的encoding  StdEncoding|URLEncoding|RawURLEncoding
	src3 := []byte("abc")
	dst3 := make([]byte, encoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(dst3, src3)
	fmt.Println("stdcoding encode:", dst3)

	//使用默认的decoding
	dst4 := make([]byte, encoding.DecodedLen(len(dst3)))
	decode, err := base64.StdEncoding.Decode(dst4, dst3)
	fmt.Println("stdenciding decode:", decode, err, dst4, string(dst4))
}
