package middleware

import (
	"20red_police/auth"
	"20red_police/common"
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
		return fmt.Errorf("serviceMethod:%s, 不在auth列表中，请配置", serviceMethod)
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

	guard, err := data.GclassTree().Pick(common.REGISTER_DATA_USER)
	if err != nil {
		return err
	}
	user, ok := guard.(data.User)
	if !ok {
		return errors.New("guard login class is err")
	}
	if _, isLogin := user.IsLogin(req.Header.BName); !isLogin {
		return errors.New("用户离线，请先登陆")
	}
	return nil
}
