package middleware

import (
	"20red_police/auth"
	"20red_police/internal/data"
	"20red_police/network"
	"20red_police/tools"
	"errors"
	"fmt"
)

func LoginGuardMiddleware(req *network.Request) error {
	return nil
	serviceMethod := req.ServiceMethod
	al, ok := auth.AuthListMapping[serviceMethod]
	if !ok {
		return fmt.Errorf("serviceMethod:%s, auth list has no this serviceMethod", serviceMethod)
	}
	if !al.IsNeedLogin {
		return nil
	}
	if err := req.CheckHeader(); err != nil {
		return err
	}
	if err := tools.Check(req.Header.Token, req.Header.BName); err != nil {
		return err
	}

	guard, err := data.MemoryUser()
	if err != nil {
		return err
	}
	user, ok := guard.(data.User)
	if !ok {
		return errors.New("guard login class is err")
	}
	if _, isLogin := user.IsLogin(req.Header.BName); !isLogin {
		return errors.New("user is offline ,please login")
	}
	return nil
}
