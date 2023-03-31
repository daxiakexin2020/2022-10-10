package protocol

import "20red_police/internal/model"

type Country struct {
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	ArchitectureNames []string `json:"architecture_names"`
}

type CountryListRequest struct {
}

type CountryListResponse struct {
	List []Country `json:"list"`
}

type FetchCountryRequest struct {
	Name string `json:"name" mapstructure:"name" validate:"required"`
}

type FetchCountryResponse struct {
	Country
}

func FormatCountryByDBToPro(model model.Country) Country {
	return Country{
		Id:                model.Id,
		Name:              model.Name,
		ArchitectureNames: model.ArchitectureNames,
	}
}
