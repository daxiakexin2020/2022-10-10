package telephone

import (
	"03/center/defined"
	"fmt"
)

type telephoneLogin struct {
	Telephone string `json:"telephone"`
	Code      string `json:"code"`
}

func NewTelephoneLogin() *telephoneLogin {
	return &telephoneLogin{}
}

func (tl *telephoneLogin) Login(req *defined.LoginRequest) error {
	fmt.Println("手机号登陆成功")
	return nil
}
