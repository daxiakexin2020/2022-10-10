package model

import "fmt"

type SymmetryEnDeEncode struct {
	Str  string   `json:"str" binding:"required"`
	Type SYM_TYPE `json:"type" binding:"required"`
}

type SymmetryEnDeDecode struct {
	Str  string   `json:"str" binding:"required"`
	Type SYM_TYPE `json:"type" binding:"required"`
}

type SYM_TYPE string

const (
	SYM_AES SYM_TYPE = "AES"
	SYM_DES          = "DES"
)

var DefalutSecretKey = []byte("123456@!,=789//.")

var SYM_MENU = map[SYM_TYPE]string{
	SYM_AES: "AES",
	SYM_DES: "DES",
}

func (st SYM_TYPE) Isvalid() bool {
	fmt.Println("st:", st)
	if _, ok := SYM_MENU[st]; !ok {
		return false
	}
	return true
}

func NewSymmetryEnDeEncode() *SymmetryEnDeEncode {
	return &SymmetryEnDeEncode{}
}
func NewSymmetryEnDeDecode() *SymmetryEnDeDecode {
	return &SymmetryEnDeDecode{}
}

type SymmetryEnDeRepo interface {
	Encode(str string, typ SYM_TYPE) (string, error)
	Decode(str string, typ SYM_TYPE) (string, error)
}
