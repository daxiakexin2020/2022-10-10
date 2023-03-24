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
	Score         uscore
	GameNumber    int
	CreateTime    string
	LastLoginTime string
	mu            sync.RWMutex `json:"-"`
}

type ulevel int

type ustatus int

type uscore int64

const (
	status_playing ustatus = iota + 1
	status_forbidden
	status_prepare
)

const level_step = 1

const (
	level_00 ulevel = iota + level_step - 1
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
	level_top = iota + 20*level_step
)

const (
	score_00 uscore = 0
	score_01        = 10000
	score_02        = 20000
	score_03        = 40000
	score_04        = 60000
	score_20        = 2<<62 - 1
)

var socre_level map[uscore]ulevel = map[uscore]ulevel{
	score_00: level_00,
	score_01: level_01,
	score_02: level_02,
	score_03: level_03,
	score_04: level_04,
	score_20: level_top,
}

var level_score map[ulevel]uscore = map[ulevel]uscore{
	level_00:  score_00,
	level_01:  score_01,
	level_02:  score_02,
	level_03:  score_03,
	level_04:  score_04,
	level_top: score_20,
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

func (u *User) nextLevel() ulevel {
	return u.Level + level_step
}

func (u *User) nextUpgradeScore() uscore {
	return level_score[u.nextLevel()]
}

func (u *User) Upgrade(addScore int64) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	currentScore := u.Score + uscore(addScore)
	nextNeedScore := u.nextUpgradeScore()
	if currentScore < nextNeedScore {
		u.Score = currentScore
		return errors.New("user score is Insufficient upgrade")
	}
	leftoverScore := currentScore - nextNeedScore
	u.Score = leftoverScore
	u.Level = u.nextLevel()
	return nil
}

func (u *User) SetForbidden() error {
	u.mu.RLock()
	defer u.mu.RUnlock()
	u.Status = status_forbidden
	return nil
}
func (u *User) SetPlaying() error {
	u.mu.RLock()
	defer u.mu.RUnlock()
	u.Status = status_playing
	return nil
}
func (u *User) SetPrepare() error {
	u.mu.RLock()
	defer u.mu.RUnlock()
	u.Status = status_prepare
	return nil
}

func (u *User) IsForbidden() bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.Status == status_forbidden
}

func (u *User) IsPrepare() bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.Status == status_prepare
}

func (u *User) IsPlaying() bool {
	u.mu.RLock()
	defer u.mu.RUnlock()
	return u.Status == status_playing
}

func (u *User) CanUpdate() error {
	if u.Id == "" || u.Name == "" || u.Pwd == "" {
		return errors.New("user field is missing")
	}
	return nil
}
