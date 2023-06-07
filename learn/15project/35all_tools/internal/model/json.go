package model

type JsonCheck struct {
	Str interface{} `json:"str" binding:"required"`
}

type JsonToGolangStruct struct {
	Str  interface{} `json:"str" binding:"required"`
	Name string      `json:"name"`
}

func NewJsonCheck() *JsonCheck {
	return &JsonCheck{}
}

func NewJsonToGolangStruct() *JsonToGolangStruct {
	return &JsonToGolangStruct{}
}

type JsonRepo interface {
	JsonCheck(model *JsonCheck) (interface{}, error)
	JsonToGolangStruct(model *JsonToGolangStruct) (interface{}, error)
}
