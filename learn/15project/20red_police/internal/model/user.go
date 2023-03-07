package model

type User struct {
	Id            string
	Name          string
	Pwd           string
	Phone         string
	Level         int
	Scorce        int64
	GamesNumber   int
	CreateTime    string
	LastLoginTime string
	Status        int
}
