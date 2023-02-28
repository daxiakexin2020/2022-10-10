package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type people struct {
	Name   string  `json:"-"`
	Age    uint8   `json:"age"`
	salary uint8   `json:"salary,omitempty"`
	Hobby  string  `json:"hobby,omitempty"`
	Jx     float64 `json:"jx"`
	IsOk   bool    `json:",omitempty"`
	B      byte    `json:"b"`
}

var _ io.Reader = (*bytes.Buffer)(nil)
var _ io.Writer = (*bytes.Buffer)(nil)

func main() {

	//编码
	v := people{
		Name:   "<zz>",
		Age:    30,
		salary: 2,
		Hobby:  "za",
		Jx:     2.2,
		B:      byte('a'),
	}
	marshal, err := json.Marshal(v)
	fmt.Println("marshal:", marshal, string(marshal), err)

	//解码
	var dst people
	err = json.Unmarshal(marshal, &dst)
	fmt.Println("unmarshal:", err, dst, dst.Hobby)

	//新建一个buffer，buffer实现了io.Writer 与io.Reader
	buf := make([]byte, 0)
	buffer := bytes.NewBuffer(buf)
	encoder := json.NewEncoder(buffer)

	//Encode将v的json编码写入输出流，并会写入一个换行符，参见Marshal函数的文档获取细节信息。
	err = encoder.Encode(v)
	fmt.Println("buffer encode:", err)

	//从buffer中获取编码后的json数据
	//dst2 := make([]byte, 2048)
	//n, err := buffer.Read(dst2)
	//fmt.Println("buffer Read:", n, err, dst2, string(dst2))

	//Decode从输入流读取下一个json编码值并保存在v指向的值里，参见Unmarshal函数的文档获取细节信息。
	var v2 people
	decoder := json.NewDecoder(buffer)
	err = decoder.Decode(&v2)
	fmt.Println("buffer decode:", v2, err)
}
