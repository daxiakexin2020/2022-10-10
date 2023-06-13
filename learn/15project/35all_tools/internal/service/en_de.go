package service

import (
	"35all_tools/internal/model"
)

type EnDeService struct {
	repo model.EnDeRepo
}

func NewEnDeService(repo model.EnDeRepo) *EnDeService {
	return &EnDeService{repo: repo}
}

func (eds *EnDeService) Md5Encode(model *model.MD5Encode) (string, error) {
	return eds.repo.MD5Encode(model.Str)
}

func (eds *EnDeService) UrlEncode(model *model.Url16Encode) (string, error) {
	return eds.repo.Url16Encode(model.Str)
}

func (eds *EnDeService) Base64Encode(model *model.Base64Encode) (string, error) {
	return eds.repo.Base64Encode(model.Str)
}

func (eds *EnDeService) Base64Decode(model *model.Base64Decode) (string, error) {
	return eds.repo.Base64Decode(model.Str)
}

func (eds *EnDeService) Escape(model *model.Escape) (string, error) {
	return eds.repo.Escape(model.Str)
}

func (eds *EnDeService) DeEscape(model *model.DeEscape) (string, error) {
	return eds.repo.DeEscape(model.Str)
}
