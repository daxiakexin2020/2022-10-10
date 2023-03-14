package server

import (
	"20red_police/protocol"
	"20red_police/tools"
	"errors"
)

/*
{"service_method":"Server.CreateRoom","meta_data":{"base":{"cookie":"1","bname":"zz"},"room_name":"r01","pmap_id":"1"}}
*/
func (s *Server) CreateRoom(req *protocol.CreateRoomRequest, res *protocol.CreateRoomResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	if err := s.check(req.Cookie, req.BName); err != nil {
		return err
	}
	_, ok := s.UserSrc.IsLogin(req.BName)
	if !ok {
		return errors.New("请先登陆")
	}
	pMap, err := s.PMapSrc.FetchPMap(req.PMapID)
	if err != nil {
		return err
	}

	room, err := s.RoomSrc.CreateRoom(req.RoomName, req.BName, pMap.Name, pMap.Count)
	if err != nil {
		return err
	}
	*res = protocol.CreateRoomResponse{protocol.FormatRoomByDBToPro(room)}
	return nil
}

/*
{"service_method":"Server.JoinRoom","meta_data":{"base":{"cookie":"1","bname":"zz"},"room_id":"r01"}}
*/
func (s *Server) JoinRoom(req *protocol.JoinRoomRequest, res *protocol.JoinRoomResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	if err := s.check(req.Cookie, req.BName); err != nil {
		return err
	}
	user, ok := s.UserSrc.IsLogin(req.BName)
	if !ok {
		return errors.New("请先登陆")
	}
	room, err := s.RoomSrc.JoinRoom(user, req.RoomID)
	if err != nil {
		return err
	}
	res = &protocol.JoinRoomResponse{protocol.FormatRoomByDBToPro(room)}
	return nil
}

/*
{"service_method":"Server.RoomList","meta_data":{"base":{"cookie":"1","bname":"zz"}}}
*/
func (s *Server) RoomList(req *protocol.RoomListRequest, res *protocol.RoomListResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	list, err := s.RoomSrc.RoomList()
	if err != nil {
		return err
	}
	*res = protocol.RoomListResponse{List: make([]protocol.Room, 0, len(list))}
	for _, room := range list {
		res.List = append(res.List, protocol.FormatRoomByDBToPro(&room))
	}
	return nil
}
