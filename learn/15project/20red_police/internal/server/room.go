package server

import (
	"20red_police/protocol"
	"20red_police/tools"
)

/*
{"service_method":"Server.CreateRoom","meta_data":{"room_name":"room02","username":"zz02", "pmap_id":"58a25e3a-c31e-11ed-3f3e-7b728b7d7fa5"}}
*/
func (s *Server) CreateRoom(req *protocol.CreateRoomRequest, res *protocol.CreateRoomResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	pMap, err := s.PMapSrc.FetchPMap(req.PMapID)
	if err != nil {
		return err
	}

	room, err := s.RoomSrc.CreateRoom(req.RoomName, req.Username, pMap.Name, pMap.Count)
	if err != nil {
		return err
	}
	*res = protocol.CreateRoomResponse{protocol.FormatRoomByDBToPro(room)}
	return nil
}

/*
{"service_method":"Server.JoinRoom","meta_data":{"username":"zz01","room_id":"66bd895e-c31e-11ed-3f3f-7b728b7d7fa5"}}
*/
func (s *Server) JoinRoom(req *protocol.JoinRoomRequest, res *protocol.JoinRoomResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	room, err := s.RoomSrc.JoinRoom(req.Username, req.RoomID)
	if err != nil {
		return err
	}
	*res = protocol.JoinRoomResponse{protocol.FormatRoomByDBToPro(room)}
	return nil
}

/*
{"service_method":"Server.RoomList","meta_data":{}}
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

/*
 */
func (s *Server) GameStart(req *protocol.GameStartRequest, res *protocol.GameStartResponse) error {
	return nil
}

/*
*
{"service_method":"Server.OutRoom","header":{"token":"1","bname":"zz"},"meta_data":{"username":"zz","room_id":"921b724c-c31d-11ed-3a67-b7d0893da24d"}}
*/
func (s *Server) OutRoom(req *protocol.OutRoomRequest, res *protocol.OutRoomResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	room, err := s.RoomSrc.OutRoom(req.Username, req.RoomID)
	if err != nil {
		return err
	}
	if room.Owner == req.Username {
		if err = s.RoomSrc.DeleteRoom(req.Username, req.RoomID); err != nil {
			return err
		}
	}
	return nil
}

/*
{"service_method":"Server.DeleteRoom","header":{"token":"1","bname":"zz"},"meta_data":{"username":"zz01","room_id":"66b63b92-c31c-11ed-3ab3-3bb14aeab3f7"}}
*/
func (s *Server) DeleteRoom(req *protocol.DelteRoomRequest, res *protocol.DeleteRoomResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	return s.RoomSrc.DeleteRoom(req.Username, req.RoomID)
}
