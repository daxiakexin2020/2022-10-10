package center

import (
	"03/center/defined"
	"03/center/telephone"
	"03/center/username"
	"fmt"
)

//代理类

type proxyLogin struct {
	il    Ilogin
	ltype Ltype
}

type Ltype string

func NewProxyLogin(ltype Ltype) *proxyLogin {
	return &proxyLogin{ltype: ltype}
}

func (pl *proxyLogin) Login(req *defined.LoginRequest) error {
	il := pl.makeIlogin()
	if il == nil {
		return fmt.Errorf("不支持此种类型的登陆策略:%v", pl.ltype)
	}
	pl.il = il
	return pl.il.Login(req)
}

func (pl *proxyLogin) ChengeLtype(ltype Ltype) {
	pl.ltype = ltype
}

func (pl *proxyLogin) makeIlogin() Ilogin {
	switch pl.ltype {
	case "username":
		return username.NewUsernameLogin()
	case "telephone":
		return telephone.NewTelephoneLogin()
	default:
		return nil
	}
}
