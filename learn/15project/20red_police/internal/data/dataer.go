package data

import "20red_police/internal/model"

type User interface {
	Register(user *model.User) error
	Login(name string, pwd string) (model.User, error)
	ForgetPwd(name string, pwd string) error
	LoginOut(name string) error
	Update(user model.User) error
	OnLineUserList() []model.User
	UserList() []model.User
}

type Room interface {
	Create(room model.Room, username string) model.Room
	Dissolve(roomID string, username string) error
	Update(room *model.Room) (model.Room, error)
	JoinRoom(player *model.Player, roomID string) error
	OutRoom(player *model.Player, roomID string) error
	List() []model.Room
}

type Player interface {
	Create(naem string) model.Player
}
