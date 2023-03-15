package protocol

import "20red_police/internal/model"

type PMap struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type CreatePMapRequest struct {
	Name  string `json:"name" mapstructure:"name" validate:"required"`
	Count int    `json:"count" mapstructure:"count" validate:"required"`
}

type CreatePMapResponse struct{ PMap }

type FetchPMapRequest struct {
	Id string `json:"id" mapstructure:"id" validate:"required"`
}

type FetchPMapResponse struct{ PMap }

type PMapListRequest struct{ Empry }

type PMapListResponse struct {
	List []PMap `json:"list"`
}

func FormatPMapByDBToPro(model *model.PMap) PMap {
	return PMap{
		Id:    model.Id,
		Name:  model.Name,
		Count: model.Count,
	}
}
