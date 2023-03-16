package server

import (
	"20red_police/protocol"
)

/*
{"service_method":"Server.CreatePMap","meta_data":{"name":"btxd","count":8}}
*/
func (s *Server) CreatePMap(req *protocol.CreatePMapRequest, res *protocol.CreatePMapResponse) error {
	pMap, err := s.PMapSrc.CreatePMap(req.Name, req.Count)
	if err != nil {
		return err
	}
	*res = protocol.CreatePMapResponse{protocol.FormatPMapByDBToPro(&pMap)}
	return nil
}

/*
{"service_method":"Server.PMapList","meta_data":{}}
*/
func (s *Server) PMapList(req *protocol.PMapListRequest, res *protocol.PMapListResponse) error {
	list, err := s.PMapSrc.PMapList()
	if err != nil {
		return err
	}
	*res = protocol.PMapListResponse{List: make([]protocol.PMap, 0)}
	for _, pmap := range list {
		res.List = append(res.List, protocol.FormatPMapByDBToPro(&pmap))
	}
	return nil
}
