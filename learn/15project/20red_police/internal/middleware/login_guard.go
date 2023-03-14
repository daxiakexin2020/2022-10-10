package middleware

import (
	"20red_police/auth"
	"20red_police/network"
	"fmt"
)

func LoginGuardMiddleware(req *network.Request) error {
	serviceMethod := req.ServiceMethod
	al, ok := auth.AuthListMapping[serviceMethod]
	if !ok {
		return fmt.Errorf("serviceMethod:%s, 不在auth列表中，请配置", serviceMethod)
	}
	if !al.IsNeedLogin {
		return nil
	}
	//todo 解耦？
	return nil
}
