package db

import "chip_database/internal/model"

type BaseInfo struct {
	db *client
}

func NewBaseInfo(db *client) *BaseInfo {
	return &BaseInfo{db: db}
}

func (bi *BaseInfo) Create(m *model.BaseInfo) error {
	return bi.db.db.Create(m).Error
}

func (bi *BaseInfo) Delete(id int) error {
	return bi.db.db.Delete(&model.BaseInfo{}, id).Error
}

func (bi *BaseInfo) Update() {

}

func (bi *BaseInfo) List() ([]*model.BaseInfo, error) {
	var baseInfo []*model.BaseInfo
	return baseInfo, bi.db.db.Find(&baseInfo).Error
}

func (bi *BaseInfo) Find(id int) (*model.BaseInfo, error) {
	var baseInfo *model.BaseInfo
	return baseInfo, bi.db.db.Find(&baseInfo, id).Error
}
