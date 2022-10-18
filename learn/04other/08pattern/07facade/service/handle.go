package service

import "errors"

type IUser interface {
	Login(phone string, code int) (*User, error)
	Register(phone string, code int) (*User, error)
}

// todo 门面接口
type IUserFacade interface {
	LoginOrRegister(phone string, code int) (*User, error)
}

type User struct {
	Phone string
}
type UserService struct {
	Users []*User
}

func (userService *UserService) Login(phone string, code int) (*User, error) {
	for _, user := range userService.Users {
		if user.Phone == phone {
			return user, nil
		}
	}
	return nil, errors.New("此账号没有注册")
}

func (userService *UserService) Register(phone string, code int) (*User, error) {
	for _, user := range userService.Users {
		if user.Phone == phone {
			return user, errors.New("此账号已经注册")
		}
	}
	user := new(User)
	user.Phone = phone
	return user, nil
}

// todo 门面
func (userService *UserService) LoginOrRegister(phone string, code int) (*User, error) {
	user, err := userService.Login(phone, code)
	if err == nil {
		return user, nil
	}
	return userService.Register(phone, code)
}

func NewUserService() *UserService {
	return &UserService{
		Users: make([]*User, 0),
	}
}
