package model

type Token struct {
	ErrMsg string `json:"errmsg"`
	ErrNo  int    `json:"errno"`
	Data   struct {
		LogId string `json:"log_id"`
	} `json:"data"`
}

func NewToken() *Token {
	return &Token{}
}

func (t *Token) ToPb() {}
