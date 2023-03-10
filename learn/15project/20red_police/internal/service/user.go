package service

import (
	"20red_police/config"
	"20red_police/internal/data"
	"20red_police/internal/model"
	"20red_police/tools"
)

type UserService struct {
	data data.User
}

func NewUserService(data data.User) *UserService {
	return &UserService{data: data}
}

func (us *UserService) Register(userModel *model.User) error {
	return us.data.Register(userModel)
}

func (us *UserService) Login(name, pwd string) (model.User, string, error) {

	umodel, err := us.data.Login(name, pwd)
	if err != nil {
		return umodel, "", err
	}

	token, err := tools.MakeToken(name, config.GetJwtConfig().TokenSecret)
	if err != nil {
		return umodel, "", err
	}

	return umodel, token, nil
}

func (us *UserService) UserList() ([]model.User, error) {
	return us.data.UserList(), nil
}
