package service

import "35all_tools/internal/model"

type ComprehensiveService struct {
	repo model.ComprehensiveRepo
}

func NewComprehensiveService(repo model.ComprehensiveRepo) *ComprehensiveService {
	return &ComprehensiveService{repo: repo}
}

func (cs *ComprehensiveService) IpInfo(model *model.IpInfo) (interface{}, error) {
	return cs.repo.IpInfo(model.Ip)
}

func (cs *ComprehensiveService) DomainMapIp(model *model.DomainMapIp) (interface{}, error) {
	return cs.repo.DomainMapIp(model.Domain)
}
