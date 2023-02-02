package user

import (
	"12lucky_draw/helper"
	"errors"
)

type viscidity int

type status int

const (
	V01 viscidity = iota + 1
	V02
	V03
	V04
	V05
	V06
	V07
	V08
	V09
	V10
)

var ViscidityPoll = []viscidity{V01, V02, V03, V04, V05, V06, V07, V08, V09, V10}
var StatusPoll = []status{forbid, normal}

const (
	forbid status = 1
	normal status = 2
)

var (
	emptyUser User
)

type User struct {
	Id         string    `json:"id"`
	Username   string    `json:"username"`
	Viscidity  viscidity `json:"viscidity"`
	status     status
	drawd      bool
	drawdLevel int
}

func (u *User) IsNormal() bool {
	return u.status == normal
}

func (u *User) IsDrawd() bool {
	return u.drawd
}

func (u *User) SetForbid() (User, error) {
	user, err := getGDB().find(u.Id)
	if err != nil {
		return emptyUser, err
	}
	user.status = forbid
	return getGDB().update(&user)
}

func (u *User) SetNormal() (User, error) {
	user, err := getGDB().find(u.Id)
	if err != nil {
		return emptyUser, err
	}
	user.status = normal
	return getGDB().update(&user)
}

func (u *User) SetDrawd(drawdLevel int) (User, error) {
	user, err := getGDB().find(u.Id)
	if err != nil {
		return emptyUser, err
	}
	if user.drawd {
		return emptyUser, errors.New("this user is drawd")
	}
	user.drawd = true
	user.drawdLevel = drawdLevel
	return getGDB().update(&user)
}

func Add(username string, vis viscidity) (User, error) {
	u := User{
		Id:        helper.Uuid(),
		Username:  username,
		Viscidity: vis,
		status:    normal,
		drawd:     false,
	}
	err := getGDB().add(&u)
	if err != nil {
		return emptyUser, err
	}
	return u, nil
}

func All() map[string]User {
	return getGDB().all()
}

func Find(uid string) (User, error) {
	user, err := getGDB().find(uid)
	if err != nil {
		return emptyUser, err
	}
	return user, nil
}

func Delete(uid string) error {
	return getGDB().delete(uid)
}

func Count() int {
	return getGDB().Count()
}
