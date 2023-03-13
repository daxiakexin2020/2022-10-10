package server

import (
	"20red_police/protocol"
	"20red_police/tools"
	"errors"
)

/*
{"service_method":"Server.Register","meta_data":{"name":"zz","pwd":"123","repwd":"123","phone":"45"}}
*/
func (s *Server) Register(req *protocol.RegisterRequest, res *protocol.RegisterResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	if req.Pwd != req.RePwd {
		return errors.New("Two passwords are different")
	}
	return s.UserSrc.Register(req.Name, req.Pwd, req.Phone)
}

/*
{"service_method":"Server.Login","meta_data":{"name":"zz","pwd":"123"}}
*/
func (s *Server) Login(req *protocol.LoginRequest, res *protocol.LoginResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	if _, ok := s.UserSrc.IsLogin(req.Name); ok {
		return nil
	}
	user, token, err := s.UserSrc.Login(req.Name, req.Pwd)
	if err != nil {
		return err
	}
	*res = protocol.LoginResponse{User: protocol.FormatUserByDBToPro(user)}
	res.Cookie = token
	res.BName = res.Name
	return nil
}

// {"service_method":"Server.LoginOut","meta_data":{"base":{"cookie":"1","bname":"zz"}}}
func (s *Server) LoginOut(req *protocol.LoginOutResquest, res *protocol.LoginOutResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	if err := s.check(req.Cookie, req.BName); err != nil {
		return err
	}
	user, ok := s.UserSrc.IsLogin(req.BName)
	if !ok {
		return nil
	}
	return s.UserSrc.LoginOut(user)
}

/*
{"service_method":"Server.UserList","meta_data":{"base":{"cookie":"1","bname":"zz"}}}
*/
func (s *Server) UserList(req *protocol.UserListRequest, res *protocol.UserListResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	if err := s.check(req.Cookie, req.BName); err != nil {
		return err
	}

	if _, ok := s.UserSrc.IsLogin(req.BName); !ok {
		return errors.New("请先登陆")
	}

	list, err := s.UserSrc.UserList()
	*res = protocol.UserListResponse{List: make([]protocol.User, 0, len(list))}
	if err != nil {
		return err
	}
	for _, data := range list {
		pro := protocol.FormatUserByDBToPro(data)
		res.List = append(res.List, pro)
	}
	return nil
}
