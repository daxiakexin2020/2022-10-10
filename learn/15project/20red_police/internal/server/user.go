package server

import (
	"20red_police/protocol"
)

/*
{"service_method":"Server.Register","meta_data":{"name":"zz02","pwd":"123","repwd":"123","phone":"45"}}
*/
func (s *Server) Register(req *protocol.RegisterRequest, res *protocol.RegisterResponse) error {
	return s.UserSrc.Register(req.Name, req.Pwd, req.RePwd, req.Phone)
}

/*
{"service_method":"Server.Login","meta_data":{"name":"zz02","pwd":"123"}}
*/
func (s *Server) Login(req *protocol.LoginRequest, res *protocol.LoginResponse) error {
	user, token, err := s.UserSrc.Login(req.Name, req.Pwd)
	if err != nil {
		return err
	}
	*res = protocol.LoginResponse{User: protocol.FormatUserByDBToPro(user)}
	res.Token = token
	res.BName = res.Name
	return nil
}

/*
{"service_method":"Server.LoginOut","header":{"token":"1","bname":"zz01"},"meta_data":{"name":"zz"}}
*/
func (s *Server) LoginOut(req *protocol.LoginOutRequest, res *protocol.LoginOutResponse) error {
	user, ok := s.UserSrc.IsLogin(req.Name)
	if !ok {
		return nil
	}
	return s.UserSrc.LoginOut(user)
}

/*
{"service_method":"Server.UserList","header":{"token":"1","bname":"out"},"meta_data":{}}
*/
func (s *Server) UserList(req *protocol.UserListRequest, res *protocol.UserListResponse) error {
	list, err := s.UserSrc.UserList()
	*res = protocol.UserListResponse{List: make([]protocol.User, 0, len(list)), Count: int64(len(list))}
	if err != nil {
		return err
	}
	for _, data := range list {
		pro := protocol.FormatUserByDBToPro(data)
		res.List = append(res.List, pro)
	}
	return nil
}
