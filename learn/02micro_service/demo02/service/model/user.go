package model

import (
	"context"
	s "service/service"
)

type User struct {
	Id    int      `json:"id"`
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby"`
}

func (u *User) GetUserInfo(ctx context.Context, client *s.UserRequest) (*s.UserResponse, error) {
	hobby := []string{"看书", "打球"}
	age := new(int32)
	*age = 1
	response := &s.UserResponse{
		Name:  "test",
		Hobby: hobby,
		Age:   age,
		Id:    client.Id,
	}
	return response, nil
}
