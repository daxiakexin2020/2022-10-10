package tools

import (
	"20red_police/components/jwt"
	"20red_police/config"
	"errors"
)

func MakeToken(consumerData interface{}, jwtSecret string) (string, error) {
	cjwt := jwt.NewToken(jwt.TokenTypeAccessToken, jwt.WithTokenSecret(jwtSecret))
	token, err := cjwt.Make(consumerData)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string, jwtSecret string) (consumerData interface{}, err error) {
	parse, err := jwt.Parse(token, jwtSecret)
	if err != nil {
		return nil, err
	}
	return parse, nil
}

func Check(token string, dest string) error {
	consumerData, err := ParseToken(token, config.GetJwtConfig().TokenSecret)
	if err != nil {
		return err
	}
	v, ok := consumerData.(*interface{})
	if !ok {
		return errors.New("consumer data is err in token ")
	}
	tname := *v
	if _, ok = tname.(string); !ok {
		return errors.New("consumer data type is error")
	}
	if tname != dest {
		return errors.New("consumer data is not eq dest!!!!!! ")
	}
	return nil
}
