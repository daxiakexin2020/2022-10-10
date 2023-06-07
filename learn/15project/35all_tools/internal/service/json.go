package service

import (
	"35all_tools/internal/model"
)

type JsonSerive struct {
	repo model.JsonRepo
}

func NewJsonService(repo model.JsonRepo) *JsonSerive {
	return &JsonSerive{repo: repo}
}

func (js *JsonSerive) JsonCheck(model *model.JsonCheck) (interface{}, error) {
	return js.repo.JsonCheck(model)
}

func (js *JsonSerive) JsonToGolangStruct(model *model.JsonToGolangStruct) (interface{}, error) {
	if model.Name == "" {
		model.Name = "RET"
	}
	return js.repo.JsonToGolangStruct(model)
}
