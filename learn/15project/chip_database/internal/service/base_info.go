package service

import (
	"chip_database/internal/data/db"
	"chip_database/internal/model"
)

type BaseInfoService struct {
	baseInfoDB *db.BaseInfo
}

func NewBaseInfoService(baseInfoDB *db.BaseInfo) *BaseInfoService {
	return &BaseInfoService{baseInfoDB: baseInfoDB}
}

func (bis *BaseInfoService) Create(m *model.BaseInfo) error {
	return bis.baseInfoDB.Create(m)
}

func (bis *BaseInfoService) Delete(id int) error {
	return bis.baseInfoDB.Delete(id)
}

func (bis *BaseInfoService) Update() {

}

func (bis *BaseInfoService) List() ([]*model.BaseInfo, error) {
	return bis.baseInfoDB.List()
}

func (bis *BaseInfoService) Find(id int) (*model.BaseInfo, error) {
	return bis.baseInfoDB.Find(id)
}
