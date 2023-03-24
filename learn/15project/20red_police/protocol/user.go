package protocol

import "20red_police/internal/model"

type Empry struct{}

type User struct {
	Id            string `json:"id"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	Level         int    `json:"level"`
	Status        int    `json:"status"`
	Score         int64  `json:"score"`
	GameNumber    int    `json:"game_number"`
	CreateTime    string `json:"create_time"`
	LastLoginTime string `json:"last_login_time"`
}

type RegisterRequest struct {
	Name  string `json:"name" mapstructure:"name" validate:"required"`
	Pwd   string `json:"pwd"  mapstructure:"pwd" validate:"required"`
	RePwd string `json:"repwd"  mapstructure:"repwd" validate:"required"`
	Phone string `json:"phone"  mapstructure:"phone" validate:"required"`
}

type RegisterResponse struct{ Empry }

type LoginRequest struct {
	Name string `json:"name" mapstructure:"name" validate:"required"`
	Pwd  string `json:"pwd"  mapstructure:"pwd" validate:"required"`
}

type LoginResponse struct {
	Header
	User
}

type LoginOutRequest struct {
	Name string `json:"name" mapstructure:"name" validate:"required"`
}

type LoginOutResponse struct{ Empry }

type UserListRequest struct{ Empry }

type UserListResponse struct {
	List  []User `json:"list"`
	Count int64  `json:"count"`
}

func FormatUserByDBToPro(model model.User) User {
	return User{
		Id:            model.Id,
		Name:          model.Name,
		Phone:         model.Phone,
		Level:         int(model.Level),
		Status:        int(model.Status),
		Score:         int64(model.Score),
		GameNumber:    model.GameNumber,
		CreateTime:    model.CreateTime,
		LastLoginTime: model.LastLoginTime,
	}
}
