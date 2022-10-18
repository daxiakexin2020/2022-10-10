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

// TODO 简单工厂模式，工厂方法主要改造这里，上游直接调用，直接new一个工厂出来，自己决定使用哪个工厂
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

// TODO 工厂方法，体现在提供出去工厂的方式，每个工厂都有自己的方法，分别提供出去，上游通过方法决定使用哪个工厂，而不是传入参数
type CreateFactory interface {
	CreateArticle() (Article, error)
}

type CreateChine struct {
}

type CreateEngine struct {
}

func (cc *CreateChine) CreateArticle() (Article, error) {
	return &ChineArticle{}, nil
}

func (ce *CreateEngine) CreateArticle() (Article, error) {
	return &EngliseArticle{Content: make([]byte, 0)}, nil
}
