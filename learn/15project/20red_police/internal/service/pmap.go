package service

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
)

type PMapService struct {
	repo data.PMap
}

func NewPMapService(pMap data.PMap) *PMapService {
	return &PMapService{repo: pMap}
}

func (pms *PMapService) FetchPMap(id string) (model.PMap, error) {
	return pms.repo.FetchPMap(id)
}
