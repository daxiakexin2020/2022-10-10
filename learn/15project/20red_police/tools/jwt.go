package tools

import "20red_police/components/jwt"

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
