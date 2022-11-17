package username

import (
	"03/center/defined"
	"fmt"
)

type usernameLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUsernameLogin() *usernameLogin {
	return &usernameLogin{}
}

func (ul *usernameLogin) Login(req *defined.LoginRequest) error {
	fmt.Println("用户名登陆成功")
	return nil
}
