package db

import (
	"chip_database/internal/model"
)

type Source struct {
	db *client
}

func NewSource(db *client) *Source {
	return &Source{db: db}
}

func (s *Source) Create(m *model.Source) error {
	return s.db.db.Create(m).Error
}

func (s *Source) Delete(id int64) error {
	m := &model.Source{Id: id}
	m.SetNotActivated()
	return s.db.db.Model(&m).Select("is_activate").Updates(m).Error
}

func (s *Source) Update(m *model.Source) error {
	return s.db.db.Model(&m).Updates(m).Error
}

func (s *Source) BatchUpdateTestId(sourceIds []int64, testId int) error {
	m := &model.Source{TestId: testId}
	m.SetActivated()
	return s.db.db.Table("source").Where("id in ? ", sourceIds).Updates(m).Error
}

func (s *Source) Find(id int64) (*model.Source, error) {
	var m *model.Source
	return m, s.db.db.Find(&m, id).Error
}
