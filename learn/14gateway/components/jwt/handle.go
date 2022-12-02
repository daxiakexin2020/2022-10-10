package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	TokenTypeAccessToken    tokenType     = "access_token"
	TokenTypeRefreshToken   tokenType     = "refresh_token"
	DefaultTokenIssuer                    = "GATEWAY"
	defaultTokenSecret                    = "ab12@!.KX2012?.Zz."
	DefalutTokenTimeoutHour time.Duration = time.Hour
	DefaultTokenTimeoutDay  time.Duration = 24 * DefalutTokenTimeoutHour
	DefaultTokenTimeoutWeek time.Duration = 7 * DefaultTokenTimeoutDay
)

var (
	ErrInvalidToken          = errors.New("invalid token")
	ErrTokenExpired          = errors.New("token expired")
	ErrTokenTypeNotSupported = errors.New("token not supported")
	ErrInvalidTokenType      = errors.New("invalid token type")
	ErrGenerateToken         = errors.New("token generate error")
)

type tokenType string

type Option func(t *token)

type claims struct {
	CustomData interface{} //用户需要参与生成token的数据
	jwt.StandardClaims
}

type token struct {
	TokenType    tokenType
	TokenSecret  string
	TokenIssuer  string
	TokenTimeout time.Duration
}

func NewToken(ttype tokenType, opts ...Option) *token {
	t := &token{
		TokenType:    ttype,
		TokenSecret:  defaultTokenSecret,
		TokenTimeout: DefaultTokenTimeoutDay,
		TokenIssuer:  DefaultTokenIssuer,
	}
	t.apply(opts)
	return t
}

func (t *token) apply(opts []Option) {
	for _, opt := range opts {
		opt(t)
	}
}

func WithTokenSecret(tokenSecret string) Option {
	return func(t *token) {
		t.TokenSecret = tokenSecret
	}
}

func WithTokenTimeout(tokenTimeout time.Duration) Option {
	return func(t *token) {
		t.TokenTimeout = tokenTimeout
	}
}

func WithTokenIssuer(tokenIssuer string) Option {
	return func(t *token) {
		t.TokenIssuer = tokenIssuer
	}
}

func (t *token) Make(customData interface{}) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims{
		customData,
		jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Add(t.TokenTimeout).Unix(),
			Issuer:    t.TokenIssuer,
		},
	})
	tokenString, err := jwtToken.SignedString([]byte(t.TokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Parse(tokenStr string, tokenSecret string) (interface{}, error) {

	stoken, err := jwt.ParseWithClaims(tokenStr, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		if ve, ok := err.(jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrInvalidTokenType
			}
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
			}
			if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrInvalidTokenType
			}
		}
		return nil, ErrInvalidToken
	}

	if !stoken.Valid {
		return nil, ErrInvalidToken
	}

	if c, ok := stoken.Claims.(*claims); ok {
		return &c.CustomData, nil
	}
	return nil, ErrInvalidToken
}
