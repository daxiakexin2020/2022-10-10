package service

import (
	"bytes"
	"errors"
	"fmt"
)

type Article interface {
	Read() string
	Write(content string)
}

type articleType string

const (
	C_TYPE articleType = "ctype"
	E_TYPE articleType = "etype"
)

type ChineArticle struct {
	Content string
}

type EngliseArticle struct {
	Content []byte
}

func (ca *ChineArticle) Read() string {
	outPut := fmt.Sprintf("%s", ca.Content)
	fmt.Println(outPut)
	return outPut
}

func (ca *ChineArticle) Write(content string) {
	var b bytes.Buffer
	b.Write([]byte(ca.Content))
	b.Write([]byte("*"))
	b.Write([]byte(content))
	ca.Content = b.String()
}
func (ea *EngliseArticle) Read() string {
	outPut := fmt.Sprintf("%s", string(ea.Content))
	fmt.Println(outPut)
	return outPut
}

func (ea *EngliseArticle) Write(content string) {
	ea.Content = append(ea.Content, []byte(content)...)
	ea.Content = append(ea.Content, []byte("\n")...)
}

func NewArticle(aType articleType) (Article, error) {
	switch aType {
	case C_TYPE:
		return &ChineArticle{}, nil
	case E_TYPE:
		return &EngliseArticle{Content: make([]byte, 0)}, nil
	default:
		return nil, errors.New("暂不支持此种方式")
	}
}
