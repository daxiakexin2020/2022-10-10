package center

import "03/center/defined"

type Ilogin interface {
	Login(req *defined.LoginRequest) error
}
