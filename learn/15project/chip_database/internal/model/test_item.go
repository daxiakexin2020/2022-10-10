package model

type TestItem struct {
	Id                      int     `json:"id"`
	BaseId                  int     `json:"base_id" form:"base_id" validate:"required"`
	Specification           string  `json:"specification" form:"specification" validate:"required"`
	ThinningTechnology      string  `json:"thinning_technology" form:"thinning_technology" validate:"required"`
	ThinningProcess         string  `json:"thinning_process" form:"thinning_process" validate:"required"`
	EncapsulationTechnology string  `json:"encapsulation_technology" form:"encapsulation_technology" validate:"required"`
	EncapsulationProcess    string  `json:"encapsulation_process" form:"encapsulation_process" validate:"required"`
	CpTestThinningRate      string  `json:"cp_test_thinning_rate" form:"cp_test_thinning_rate" validate:"required"`
	CpTestThinningThickness string  `json:"cp_test_thinning_thickness" form:"cp_test_thinning_thickness" validate:"required"`
	CpTestMap               string  `json:"cp_test_map" form:"cp_test_map" validate:"required"`
	SourceIds               []int64 `json:"source_ids" form:"source_ids" gorm:"-"`
}

type DeleteTestItem struct {
	Id int `json:"id" form:"id" validate:"required"`
}

type ListTestItem struct {
	BaseId int `json:"base_id" form:"base_id" validate:"required"`
}

type FindTestItem struct {
	Id int `json:"id" form:"id" validate:"required"`
}
