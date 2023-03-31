package server

import "20red_police/protocol"

/*
*
{"service_method":"Server.FetchCountry","meta_data":{"name":"sl"}}
*/
func (s *Server) FetchCountry(req *protocol.FetchCountryRequest, res *protocol.FetchCountryResponse) error {
	country, err := s.CountrySrc.FetchCountry(req.Name)
	if err != nil {
		return err
	}
	*res = protocol.FetchCountryResponse{protocol.FormatCountryByDBToPro(country)}
	return nil
}

/*
{"service_method":"Server.CountryList"}
*/
func (s *Server) CountryList(req *protocol.CountryListRequest, res *protocol.CountryListResponse) error {
	list := s.CountrySrc.CountryList()
	*res = protocol.CountryListResponse{List: make([]protocol.Country, 0, len(list))}
	for _, model := range list {
		res.List = append(res.List, protocol.FormatCountryByDBToPro(model))
	}
	return nil
}
