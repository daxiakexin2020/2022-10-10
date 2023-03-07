package service

import "20red_police/internal/data"

type UserService struct {
	data data.User
}

func NewUserService(data data.User) *UserService {
	return &UserService{data: data}
}
