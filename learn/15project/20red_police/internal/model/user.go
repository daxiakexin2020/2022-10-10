package model

import "20red_police/tools"

type User struct {
	Id            string
	Name          string
	Pwd           string
	Phone         string
	Level         ulevel
	Status        ustatus
	Scorce        int64
	GamesNumber   int
	CreateTime    string
	LastLoginTime string
}

type ulevel int

type ustatus int

const (
	level_01 ulevel = iota + 1
	level_02
	level_03
	level_04
	level_05
	level_06
	level_07
	level_08
	level_09
	level_10
)

const (
	status_normal ustatus = iota + 1
	status_forbidden
)

func NewUserModel(name, pwd, phone string) *User {
	return &User{
		Id:         tools.UUID(),
		Name:       name,
		Pwd:        pwd,
		Phone:      phone,
		Level:      level_01,
		Status:     status_normal,
		CreateTime: tools.NowTimeFormatTimeToString(),
	}
}
