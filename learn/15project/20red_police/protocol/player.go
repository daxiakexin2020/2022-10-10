package protocol

import "20red_police/internal/model"

type Player struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CountryName string `json:"country_name"`
	Color       string `json:"color"`
	Status      bool   `json:"status"`
	OutCome     bool   `json:"out_come"` //结局
}

func FormatPlayerByDBToPro(model *model.Player) *Player {
	return &Player{
		Id:          model.Id,
		Name:        model.Name,
		CountryName: model.CountryName,
		Color:       model.Color,
		Status:      model.Status,
		OutCome:     model.OutCome,
	}
}
