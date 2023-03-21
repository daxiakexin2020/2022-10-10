package service

import (
	"20red_police/config"
	"20red_police/internal/data"
	"20red_police/internal/model"
	"20red_police/tools"
	"errors"
)

type UserService struct {
	data data.User
}

func NewUserService(data data.User) *UserService {
	return &UserService{data: data}
}

func (us *UserService) Register(name, pwd, repwd, phone string) error {
	if pwd != repwd {
		return errors.New("Two passwords are different")
	}
	userModel := model.NewUserModel(name, pwd, phone)
	err := us.data.Register(userModel)
	if err == nil {
		//stores.UserMemorySyncFile(userModel)
	}
	return err
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

func (us *UserService) IsLogin(name string) (model.User, bool) {
	return us.data.IsLogin(name)
}

func (us *UserService) LoginOut(user model.User) error {
	return us.data.LoginOut(user)
}

func (us *UserService) UserList() ([]model.User, error) {
	return us.data.UserList(), nil
}
