package server

import (
	"20red_police/protocol"
)

/*
{"service_method":"Server.CreateRoom","meta_data":{"room_name":"room01","username":"zz", "pmap_id":"5632f972-c4a2-11ed-25f0-dbf3c9c73d0c"}}
*/
func (s *Server) CreateRoom(req *protocol.CreateRoomRequest, res *protocol.CreateRoomResponse) error {
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
{"service_method":"Server.JoinRoom","meta_data":{"username":"zz","room_id":"bf56f3ae-c49d-11ed-2b99-3bc084b0fb1c"}}
*/
func (s *Server) JoinRoom(req *protocol.JoinRoomRequest, res *protocol.JoinRoomResponse) error {
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
{"service_method":"Server.GameStart","header":{"token":"1","bname":"zz"},"meta_data":{"username":"zz","room_id":"32da7c5c-c3a3-11ed-3bc2-8fbe2ae760fe"}}
*/
func (s *Server) GameStart(req *protocol.GameStartRequest, res *protocol.GameStartResponse) error {
	if err := s.RoomSrc.GameStart(req.Username, req.RoomID); err != nil {
		return err
	}
	return nil
}

/*
{"service_method":"Server.OutRoom","header":{"token":"1","bname":"zz"},"meta_data":{"username":"zz","room_id":"32da7c5c-c3a3-11ed-3bc2-8fbe2ae760fe"}}
*/
func (s *Server) OutRoom(req *protocol.OutRoomRequest, res *protocol.OutRoomResponse) error {
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
	return s.RoomSrc.DeleteRoom(req.Username, req.RoomID)
}

/*
{"service_method":"Server.UpdateRoomPlayer","header":{"token":"1","bname":"zz"},"meta_data":{"username":"zz01","room_id":"66b63b92-c31c-11ed-3ab3-3bb14aeab3f7","status":true}}
*/
func (s *Server) UpdateRoomPlayer(req *protocol.UpdateRoomPlayerRequest, res *protocol.UpdateRoomPlayerResponse) error {
	return s.RoomSrc.UpdateRoomPlayer(req.RoomID, req.Username, req.Status)
}

func (s *Server) Kick(req *protocol.KickRequest, res *protocol.KickResponse) error {
	return nil
}
