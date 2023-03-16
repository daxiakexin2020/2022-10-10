package middleware

import (
	"20red_police/network"
	"20red_police/tools"
)

func ValidatorMiddleWare(request *network.Request) error {
	if err := tools.Validator(request.MethodReflectData); err != nil {
		return err
	}
	return nil
}
