package server

import (
	"20red_police/internal/model"
	"20red_police/protocol"
	"20red_police/tools"
	"errors"
)

// {"service_method":"Server.Register","meta_data":{"name":"zz","pwd":"123","repwd":"123","phone":"45"}}
func (s *Server) Register(req *protocol.RegisterRequest, res *protocol.RegisterResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	if req.Pwd != req.RePwd {
		return errors.New("Two passwords are different")
	}
	userModel := model.NewUserModel(req.Name, req.Pwd, req.Phone)
	return s.UserSrc.Register(userModel)
}

// {"service_method":"Server.Login","meta_data":{"name":"zz","pwd":"123"}}
func (s *Server) Login(req *protocol.LoginRequest, res *protocol.LoginResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
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

// {"service_method":"Server.UserList","meta_data":{"base":{"cookie":"1","bname":"zz"}}}
func (s *Server) UserList(req *protocol.UserListRequest, res *protocol.UserListResponse) error {
	if err := tools.Validator(req); err != nil {
		return err
	}
	if err := s.check(req.Cookie, req.BName); err != nil {
		return err
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
