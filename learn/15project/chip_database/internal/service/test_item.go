package service

import (
	"chip_database/internal/data/db"
	"chip_database/internal/model"
)

type TestItemService struct {
	testItemDB *db.TestItem
}

func NewTestItemService(testItemDB *db.TestItem) *TestItemService {
	return &TestItemService{testItemDB: testItemDB}
}

func (tis *TestItemService) Create(m *model.TestItem) error {
	return tis.testItemDB.Create(m)
}

func (tis *TestItemService) Delete(id int) error {
	return tis.testItemDB.Delete(id)
}

func (tis *TestItemService) Update(m *model.TestItem) error {
	return tis.testItemDB.Update(m)
}

func (tis *TestItemService) Find(id int) (*model.TestItem, error) {
	return tis.testItemDB.Find(id)
}

func (tis *TestItemService) List() ([]*model.TestItem, error) {
	return tis.testItemDB.List()
}

func (tis *TestItemService) GetAllByBaseId(baseId int) ([]*model.TestItem, error) {
	return tis.testItemDB.GetAllByBaseId(baseId)
}
