package model

type Base struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func (b *Base) IsSuccess() bool {
	return b.Code == "0"
}
