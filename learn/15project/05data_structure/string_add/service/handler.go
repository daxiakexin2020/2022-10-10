package service

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		tn := rand.Intn(len(letterBytes))
		b[i] = letterBytes[tn]
	}
	return string(b)
}

// 第一种 +
func plusConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

// 第二种  fmt.Sprintf
func sprintfConcat(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

// 第三种 strings.Builder
func builderConcat(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

// 第四种 bytes.Buffer
func bufferConcat(n int, str string) string {
	buf := new(bytes.Buffer)
	for j := 0; j < n; j++ {
		buf.WriteString(str)
	}
	return buf.String()
}

// 第五种 []byte
func byteConcat(n int, str string) string {
	buf := make([]byte, 0)
	for j := 0; j < n; j++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

// 第六种 改进[]byte，加容量
func preByteConcat(n int, str string) string {
	buf := make([]byte, 0, n*len(str))
	for j := 0; j < n; j++ {
		buf = append(buf, str...)
	}
	return string(buf)
}
