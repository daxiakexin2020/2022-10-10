package bus

import (
	"errors"
	"redis/service"
)

const (
	Auth_ID        = "id"
	Auth_Name      = "name"
	Auth_Token     = "token"
	Auth_Namespace = "namesapce"
	Auth_Telephone = "telephone"
)

type Auth struct {
	Id        string
	Name      string
	Token     string
	Namespace string
	Telephone string
	err       error `json:"-"`
}

type Option func(a *Auth)

func WithNamesapce(namespace string) Option {
	return func(a *Auth) {
		a.Namespace = namespace
	}
}

func WithTelephone(telephone string) Option {
	return func(a *Auth) {
		a.Telephone = telephone
	}
}

func (a *Auth) apply(opts ...Option) {
	for _, opt := range opts {
		opt(a)
	}
}

func NewAuth(id, name, token string, opts ...Option) *Auth {
	a := &Auth{Id: id, Name: name, Token: token}
	a.apply(opts...)
	return a
}

func (a *Auth) isValid() bool {
	if a.Id == "" || a.Name == "" || a.Token == "" {
		return false
	}
	return true
}

func (a *Auth) Set(client *service.Rclient) error {
	if !a.isValid() {
		return errors.New("auth field is missing")
	}
	a.set(client, a.Token, Auth_ID, a.Id)
	a.set(client, a.Token, Auth_Name, a.Name)
	a.set(client, a.Token, Auth_Token, a.Token)
	a.set(client, a.Token, Auth_Namespace, a.Namespace)
	a.set(client, a.Token, Auth_Telephone, a.Telephone)
	return a.err
}

func (a *Auth) set(client *service.Rclient, key, field string, val interface{}) {

	if client == nil {
		a.err = errors.New("redis client is nil")
		return
	}
	if a.err != nil {
		return
	}
	_, err := client.HSet(key, field, val)
	a.err = err
}
