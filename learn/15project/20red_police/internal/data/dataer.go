package data

import "20red_police/internal/model"

type User interface {
	Class
	Register(user *model.User) error
	Login(name string, pwd string) (model.User, error)
	IsLogin(name string) (model.User, bool)
	ForgetPwd(name string, pwd string) error
	LoginOut(user model.User) error
	Update(user model.User) (model.User, error)
	OnLineUserList() []model.User
	UserList() []model.User
	IsOnLine(name string) (model.User, bool)
	FetchUser(name string) (model.User, error)
	UserCanTransformPlayer(name string) (model.User, error)
}

type Room interface {
	Class
	Create(room *model.Room) (model.Room, error)
	Dissolve(roomID string, username string) error
	Update(room *model.Room) (model.Room, error)
	JoinRoom(player *model.Player, roomID string) (model.Room, error)
	OutRoom(playerName string, roomID string) error
	RoomList() []model.Room
	FetchRoom(roomID string) (model.Room, error)
	DeleteRoom(playerName string, roomID string) error
	GameStart(username, roomID string) error
	UpdateRoomPlayer(roomID string, username string, status bool) error
	Broadcast(rootID string) error
}

type Player interface {
	Create(model model.Player) model.Player
}

type PMap interface {
	Class
	Create(pmap *model.PMap) (model.PMap, error)
	List() []model.PMap
	FetchPMap(id string) (model.PMap, error)
}

type Country interface {
	Class
	CountryList() []model.Country
	FetchCountry(name string) (model.Country, error)
}

type Architecture interface {
	ArchitectureList() []model.Architecture
	FetchArchitecture(Name string) (model.Architecture, error)
}

type Arm interface {
	IsMineCar(name string) bool
	ArmList() []model.Arm
	FetchArm(Name string) (model.Arm, error)
}

type Class interface {
	Name() string
}
