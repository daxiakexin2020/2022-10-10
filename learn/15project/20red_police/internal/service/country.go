package service

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
)

type CountryService struct {
	countryRepo data.Country
}

func NewCountryService(countryRepo data.Country) *CountryService {
	return &CountryService{countryRepo: countryRepo}
}

func (c *CountryService) CountryList() []model.Country {
	return c.countryRepo.CountryList()
}

func (c *CountryService) FetchCountry(name string) (model.Country, error) {
	return c.countryRepo.FetchCountry(name)
}
