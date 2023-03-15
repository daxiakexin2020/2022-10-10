package model

import (
	"20red_police/tools"
	"errors"
	"sync"
)

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
	Mu            sync.RWMutex `json:"-"`
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
	status_playing ustatus = iota + 1
	status_forbidden
	status_prepare
)

func NewUserModel(name, pwd, phone string) *User {
	return &User{
		Id:         tools.UUID(),
		Name:       name,
		Pwd:        pwd,
		Phone:      phone,
		Level:      level_01,
		Status:     status_prepare,
		CreateTime: tools.NowTimeFormatTimeToString(),
	}
}

func (u *User) SetForbidden() error {
	u.Mu.RLock()
	defer u.Mu.RUnlock()
	u.Status = status_forbidden
	return nil
}
func (u *User) SetPlaying() error {
	u.Mu.RLock()
	defer u.Mu.RUnlock()
	u.Status = status_playing
	return nil
}
func (u *User) SetPrepare() error {
	u.Mu.RLock()
	defer u.Mu.RUnlock()
	u.Status = status_prepare
	return nil
}

func (u *User) IsForbidden() bool {
	u.Mu.RLock()
	defer u.Mu.RUnlock()
	return u.Status == status_forbidden
}

func (u *User) IsPrepare() bool {
	u.Mu.RLock()
	defer u.Mu.RUnlock()
	return u.Status == status_prepare
}

func (u *User) IsPlaying() bool {
	u.Mu.RLock()
	defer u.Mu.RUnlock()
	return u.Status == status_playing
}

func (u *User) CanUpdate() error {
	if u.Id == "" || u.Name == "" || u.Pwd == "" {
		return errors.New("user field is missing")
	}
	return nil
}
