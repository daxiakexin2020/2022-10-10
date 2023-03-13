package protocol

import "20red_police/internal/model"

type Empry struct{}

type User struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Level         int    `json:"level"`
	Status        int    `json:"status"`
	Scorce        int64  `json:"scorce"`
	GamesNumber   int    `json:"games_number"`
	CreateTime    string `json:"create_time"`
	LastLoginTime string `json:"last_login_time"`
}

type RegisterRequest struct {
	Name  string `json:"name" mapstructure:"name" validate:"required"`
	Pwd   string `json:"pwd"  mapstructure:"pwd" validate:"required"`
	RePwd string `json:"pwd"  mapstructure:"repwd" validate:"required"`
	Phone string `json:"phone"  mapstructure:"phone" validate:"required"`
}

type RegisterResponse struct {
	Empry
}

type UserListRequest struct {
	Base
}

type UserListResponse struct {
	List []User `json:"list"`
}

type LoginRequest struct {
	Name string `json:"name" mapstructure:"name" validate:"required"`
	Pwd  string `json:"pwd"  mapstructure:"pwd" validate:"required"`
}

type LoginResponse struct {
	Base
	User
}

type LoginOutResquest struct {
	Base
}

type LoginOutResponse struct {
	Empry
}

func FormatUserByDBToPro(model model.User) User {
	return User{
		Id:            model.Id,
		Name:          model.Name,
		Phone:         model.Phone,
		Level:         int(model.Level),
		Status:        int(model.Status),
		GamesNumber:   model.GamesNumber,
		CreateTime:    model.CreateTime,
		LastLoginTime: model.LastLoginTime,
	}
}
