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
	GameNumber    int
	CreateTime    string
	LastLoginTime string
	Mu            sync.RWMutex `json:"-"`
}

type ulevel int

type ustatus int

type usocre int64

const (
	status_playing ustatus = iota + 1
	status_forbidden
	status_prepare
)

const (
	level_00 ulevel = iota
	level_01
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
	score_00 usocre = 0
	score_01        = 10000
	score_02        = 20000
	score_03        = 40000
	score_04        = 60000
)

var socre_level map[usocre]ulevel = map[usocre]ulevel{
	score_00: level_00,
	score_01: level_01,
	score_02: level_02,
	score_03: level_03,
	score_04: level_04,
}

var level_score map[ulevel]usocre = map[ulevel]usocre{
	level_00: score_00,
	level_01: score_01,
	level_02: score_02,
	level_03: score_03,
	level_04: score_04,
}

//chan username score

type Result struct {
	Username string
	Score    int64
}

func NewUserModel(name, pwd, phone string) *User {
	return &User{
		Id:         tools.UUID(),
		Name:       name,
		Pwd:        pwd,
		Phone:      phone,
		Level:      level_00,
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
