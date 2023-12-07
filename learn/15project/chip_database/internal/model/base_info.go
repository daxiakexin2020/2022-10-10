package model

type BaseInfo struct {
	ID                     int    `json:"id"`
	IntegratedCircuitClass string `json:"integrated_circuit_class" form:"integrated_circuit_class" validate:"required"`
	Technology             string `json:"technology" form:"technology"  validate:"required"`
	ManufacturingProcess   string `json:"manufacturing_process" form:"manufacturing_process"  validate:"required"`
	ChipType               string `json:"chip_type" form:"chip_type"  validate:"required"`
	ModelNumber            string `json:"model_number" form:"model_number"  validate:"required"`
	ChipSize               string `json:"chip_size" form:"chip_size"  validate:"required"`
	DesignCompany          string `json:"design_company" form:"design_company"  validate:"required"`
	FilmStreamingPlant     string `json:"film_streaming_plant" form:"film_streaming_plant"  validate:"required"`
	WaferTestFactory       string `json:"wafer_test_factory" form:"wafer_test_factory"  validate:"required"`
}

type Tree struct {
	ID                     int         `json:"id"`
	IntegratedCircuitClass string      `json:"integrated_circuit_class" `
	Technology             string      `json:"technology"`
	ManufacturingProcess   string      `json:"manufacturing_process"`
	ChipType               string      `json:"chip_type"`
	ModelNumber            string      `json:"model_number"`
	ChipSize               string      `json:"chip_size"`
	DesignCompany          string      `json:"design_company"`
	FilmStreamingPlant     string      `json:"film_streaming_plant"`
	WaferTestFactory       string      `json:"wafer_test_factory"`
	Children               []*TestItem `json:"children"`
}

type DeleteBaseInfo struct {
	Id int `json:"id" form:"id" validate:"required"`
}

type FindBaseInfo struct {
	Id int `json:"id" form:"id" validate:"required"`
}

func (bi *BaseInfo) GenerateTree() *Tree {
	return &Tree{
		ID:                     bi.ID,
		IntegratedCircuitClass: bi.IntegratedCircuitClass,
		Technology:             bi.Technology,
		ManufacturingProcess:   bi.ManufacturingProcess,
		ChipType:               bi.ChipType,
		ModelNumber:            bi.ModelNumber,
		ChipSize:               bi.ChipSize,
		DesignCompany:          bi.DesignCompany,
		FilmStreamingPlant:     bi.FilmStreamingPlant,
		WaferTestFactory:       bi.WaferTestFactory,
		Children:               make([]*TestItem, 0),
	}
}
