package server

import (
	"20red_police/protocol"
	"20red_police/tools"
)

/*
{"service_method":"Server.CreatePMap","meta_data":{"base":{"cookie":"1","bname":"zz"},"name":"1","count":8}}
*/
func (s *Server) CreatePMap(req *protocol.CreatePMapRequest, res *protocol.CreatePMapResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	pMap, err := s.PMapSrc.CreatePMap(req.Name, req.Count)
	if err != nil {
		return err
	}
	*res = protocol.CreatePMapResponse{protocol.FormatPMapByDBToPro(&pMap)}
	return nil
}

/*
{"service_method":"Server.PMapList","meta_data":{"base":{"cookie":"1","bname":"zz"},"name":"1","count":8}}
*/
func (s *Server) PMapList(req *protocol.PMapListRequest, res *protocol.PMapListResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
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
