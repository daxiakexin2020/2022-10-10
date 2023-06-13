package model

type MD5Encode struct {
	Str string `json:"str" binding:"required"`
}

type Url16Encode struct {
	Str string `json:"str" binding:"required"`
}

type Base64Encode struct {
	Str string `json:"str" binding:"required"`
}

type Base64Decode struct {
	Str string `json:"str" binding:"required"`
}

type Escape struct {
	Str string `json:"str" binding:"required"`
}

type DeEscape struct {
	Str string `json:"str" binding:"required"`
}

func NewMD5Encode() *MD5Encode {
	return &MD5Encode{}
}

func NewUrl16Encode() *Url16Encode {
	return &Url16Encode{}
}

func NewBase64Encode() *Base64Encode {
	return &Base64Encode{}
}

func NewBase64Decode() *Base64Decode {
	return &Base64Decode{}
}

func NewEscape() *Escape {
	return &Escape{}
}

func NewDeEscape() *DeEscape {
	return &DeEscape{}
}

type EnDeRepo interface {
	MD5Encode(str string) (string, error)
	Url16Encode(str string) (string, error)
	Base64Encode(str string) (string, error)
	Base64Decode(str string) (string, error)
	Escape(str string) (string, error)
	DeEscape(str string) (string, error)
}
