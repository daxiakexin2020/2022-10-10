package protocol

import "20red_police/internal/model"

type Player struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	CountryName string `json:"country_name"`
	Color       string `json:"color"`
	Money       int32
	Status      bool `json:"status"`
	OutCome     bool `json:"out_come"` //结局
}

type BuildArchitectureRequest struct {
	Username string `json:"username" mapstructure:"username" validate:"required"`
	RoomID   string `json:"room_id" mapstructure:"room_id" validate:"required"`
	ARName   string `json:"ar_name" mapstructure:"ar_name" validate:"required"`
}

type BuildArchitectureResponse struct{}

type BuildArmRequest struct {
	Username string `json:"username" mapstructure:"username" validate:"required"`
	RoomID   string `json:"room_id" mapstructure:"room_id" validate:"required"`
	ArmName  string `json:"arm_name" mapstructure:"arm_name" validate:"required"`
}

type BuildArmResponse struct{}

func FormatPlayerByDBToPro(model *model.Player) *Player {
	return &Player{
		Id:          model.Id,
		Name:        model.Name,
		CountryName: model.CountryName,
		Color:       model.Color,
		Status:      model.Status,
		OutCome:     model.OutCome,
		Money:       model.Money,
	}
}
