package server

import (
	"20red_police/protocol"
	"errors"
	"fmt"
)

/*
{"service_method":"Server.CreateRoom","meta_data":{"room_name":"room01","username":"zz01", "pmap_id":"981e3bfe-cfaf-11ed-3930-1726d5e31d44","init_price":100}}
*/
func (s *Server) CreateRoom(req *protocol.CreateRoomRequest, res *protocol.CreateRoomResponse) error {
	//pMap, err := s.PMapSrc.FetchPMap(req.PMapID)
	//if err != nil {
	//	return err
	//}
	//
	//room, err := s.RoomSrc.CreateRoom(req.RoomName, req.Username, pMap.Name, pMap.Count)
	//if err != nil {
	//	return err
	//}
	//*res = protocol.CreateRoomResponse{protocol.FormatRoomByDBToPro(room)}
	//return nil

	room, err := s.RoomSrc.CreateRoom(req.RoomName, req.Username, "test", 8, req.InitPrice)
	if err != nil {
		return err
	}
	*res = protocol.CreateRoomResponse{protocol.FormatRoomByDBToPro(room)}
	return nil

}

/*
{"service_method":"Server.JoinRoom","meta_data":{"username":"zz02","room_id":"1"}}
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

/**
{"service_method":"Server.GameOver","header":{"token":"1","bname":"zz"},"meta_data":{"username":"zz","room_id":"1"}}
*/

func (s *Server) GameOver(req *protocol.GameOverRequest, res *protocol.GameOverResponse) error {
	if err := s.RoomSrc.GameOver(req.RoomID); err != nil {
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

/*
{"service_method":"Server.Kick","header":{"token":"1","bname":"zz"},"meta_data":{"username":"zz01","room_id":"1","bekick_username":"zz02"}}
*/
func (s *Server) Kick(req *protocol.KickRequest, res *protocol.KickResponse) error {
	if err := s.RoomSrc.Kick(req.Username, req.BekickedUsername, req.RoomID); err != nil {
		return err
	}
	return nil
}

/*
{"service_method":"Server.Broadcast","header":{"token":"1","bname":"zz"},"meta_data":{"username":"zz01","room_id":"cf25eb94-c882-11ed-2ada-9f2dd5d7862e"}}
*/
func (s *Server) Broadcast(req *protocol.BroadcastRequest, res *protocol.BroadcastResponse) error {
	user, err := s.UserSrc.FetchUser(req.Username)
	if err != nil {
		return err
	}
	room, err := s.RoomSrc.FetchRoom(req.RoomID)
	if err != nil {
		return err
	}
	if user.Name != room.Owner {
		return errors.New("room owner can Broadcast only")
	}
	data := fmt.Sprintf("room:[%s]，room ID:[%s]，Invite you to play.... ", room.Name, room.Id)
	s.NetworkSrc.Broadcast(data)
	return nil
}

/*
{"service_method":"Server.BuildArchitecture","meta_data":{"username":"zz01","room_id":"1","ar_name":"sjby"}}
*/
func (s *Server) BuildArchitecture(req *protocol.BuildArchitectureRequest, res *protocol.BuildArchitectureResponse) error {
	room, err := s.RoomSrc.FetchRoom(req.RoomID)
	if err != nil {
		return err
	}
	if !room.IsPlaying() {
		return errors.New("game is not start")
	}
	player, err := room.FetchPlayer(req.Username)
	if err != nil {
		return err
	}
	return s.PlayerSrc.BuildArchitecture(player, req.RoomID, req.ARName)
}

/*
{"service_method":"Server.BuildArm","meta_data":{"username":"zz01","room_id":"1","arm_name":"mjkc"}}
*/
func (s *Server) BuildArm(req *protocol.BuildArmRequest, res *protocol.BuildArmResponse) error {
	room, err := s.RoomSrc.FetchRoom(req.RoomID)
	if err != nil {
		return err
	}
	if !room.IsPlaying() {
		return errors.New("game is not start")
	}
	player, err := room.FetchPlayer(req.Username)
	if err != nil {
		return err
	}
	return s.PlayerSrc.BuildArm(player, req.RoomID, req.ArmName)
}
