package main

import (
	"fmt"
	"net/url"
)

func main() {
	u := "https://www.baidu.com"

	//QueryEscape函数对s进行转码使之可以安全的用在URL查询里。
	escape := url.QueryEscape(u)
	fmt.Println("escape:", escape)

	//QueryUnescape函数用于将QueryEscape转码的字符串还原。它会把%AB改为字节0xAB，将'+'改为' '。如果有某个%后面未跟两个十六进制数字，本函数会返回错误。
	unescape, err := url.QueryUnescape(escape)
	if err != nil {
		panic(err)
	}
	fmt.Println("unescape:", unescape)
}
