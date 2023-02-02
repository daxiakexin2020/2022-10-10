package server

import (
	muser "12lucky_draw/model/user"
	"errors"
	"strconv"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (us *UserService) Add(username string) (muser.User, error) {
	index := muser.Count() % len(muser.ViscidityPoll)
	username = username + "_" + strconv.Itoa(index)
	viscidity := muser.ViscidityPoll[index]
	user, err := muser.Add(username, viscidity)
	return user, err
}

func (us *UserService) Delete(uid string) error {
	return muser.Delete(uid)
}

func (us *UserService) All() map[string]muser.User {
	return muser.All()
}

func (us *UserService) Find(uid string) (muser.User, error) {
	return muser.Find(uid)
}

func (us *UserService) Draw(uid string) (int, error) {
	user, err := us.Find(uid)
	if err != nil {
		return 0, err
	}
	if !user.IsNormal() {
		return 0, errors.New("this user is forbidden")
	}

	if user.IsDrawd() {
		return 0, errors.New("you arr drawd!!!")
	}
	ds := NewDrawService()
	return ds.Draw(int32(user.Viscidity))
}
