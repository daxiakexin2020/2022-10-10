package data

import "20red_police/internal/model"

type User interface {
	Register(user *model.User) error
	Login(name string, pwd string) (model.User, error)
	IsLogin(name string) (model.User, bool)
	ForgetPwd(name string, pwd string) error
	LoginOut(user model.User) error
	Update(user model.User) error
	OnLineUserList() []model.User
	UserList() []model.User
	IsOnLine(name string) (model.User, bool)
	ClassR
}

type Room interface {
	Create(room *model.Room) (model.Room, error)
	Dissolve(roomID string, username string) error
	Update(room *model.Room) (model.Room, error)
	JoinRoom(player *model.Player, roomID string) (model.Room, error)
	OutRoom(player *model.Player, roomID string) error
	List() []model.Room
	Broadcast(rootID string) error
}

type Player interface {
	Create(model model.Player) model.Player
}

type PMap interface {
	Create(pmap *model.PMap) (model.PMap, error)
	List() []model.PMap
	FetchPMap(id string) (model.PMap, error)
}

type ClassR interface {
	Name() string
}
