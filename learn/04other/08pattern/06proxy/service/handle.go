package service

import "fmt"

type IUser interface {
	login(username, passwd string) error
}

type User struct{}

func NewUser() *User {
	return &User{}
}

type UserProxy struct {
	user IUser
}

func (u *User) login(username, passwd string) error {
	fmt.Println("user 开始处理··········")
	return nil
}

func NewUserProxy(user IUser) *UserProxy {
	return &UserProxy{user: user}
}

func (up *UserProxy) Login(username, passwd string) error {
	//todo 做一些代理做的功能，例如统一的验证....
	return up.user.login(username, passwd)
}
