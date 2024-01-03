package db

import "chip_database/internal/model"

type TestItem struct {
	db *client
}

func NewTestItem(db *client) *TestItem {
	return &TestItem{db: db}
}

func (ti *TestItem) Create(m *model.TestItem) error {
	return ti.db.db.Create(m).Error
}

func (ti *TestItem) Delete(id int) error {
	return ti.db.db.Delete(&model.TestItem{}, id).Error
}

func (ti *TestItem) Update(m *model.TestItem) error {
	return ti.db.db.Model(&m).Updates(m).Error
}

func (ti *TestItem) Find(id int) (*model.TestItem, error) {
	var testItem *model.TestItem
	return testItem, ti.db.db.Find(&testItem, id).Error
}

func (ti *TestItem) GetAllByBaseId(baseId int) ([]*model.TestItem, error) {
	items := make([]*model.TestItem, 0)
	err := ti.db.db.Where("base_id = ?", baseId).Find(&items).Error
	return items, err
}

func (ti *TestItem) List() ([]*model.TestItem, error) {
	var items []*model.TestItem
	return items, ti.db.db.Find(&items).Error
}
