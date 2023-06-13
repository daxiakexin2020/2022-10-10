package service

import (
	"35all_tools/internal/model"
	"errors"
)

type SymmetryEnDeService struct {
	repo model.SymmetryEnDeRepo
}

func NewSymmetryEnDeService(repo model.SymmetryEnDeRepo) *SymmetryEnDeService {
	return &SymmetryEnDeService{repo: repo}
}

func (seds *SymmetryEnDeService) Encode(model *model.SymmetryEnDeEncode) (string, error) {
	if !model.Type.Isvalid() {
		return "", errors.New("暂不支持此种加密方式：" + string(model.Type))
	}
	return seds.repo.Encode(model.Str, model.Type)
}

func (seds *SymmetryEnDeService) Decode(model *model.SymmetryEnDeDecode) (string, error) {
	if !model.Type.Isvalid() {
		return "", errors.New("暂不支持此种加密方式：" + string(model.Type))
	}
	return seds.repo.Decode(model.Str, model.Type)
}
