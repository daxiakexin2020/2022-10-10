package memory

import (
	"20red_police/common"
	"20red_police/internal/data"
	"20red_police/internal/model"
	"20red_police/tools"
	"errors"
	"fmt"
	"sync"
)

type User struct {
}

var _ data.User = (*User)(nil)

type users struct {
	list  map[string]*model.User
	mu    sync.RWMutex
	locks map[string]*sync.Mutex
}

type onLineUsers struct {
	list map[string]*model.User
	mu   sync.RWMutex
}

var (
	gonce        sync.Once
	gusers       *users
	gonLineUsers *onLineUsers
	emptyUser    model.User
)

func init() {
	gonce = sync.Once{}
	gonce.Do(func() {
		gusers = &users{
			list:  make(map[string]*model.User),
			locks: map[string]*sync.Mutex{},
		}
		gonLineUsers = &onLineUsers{
			list: make(map[string]*model.User),
		}
	})
}

func NewUser() data.User {
	u := &User{}
	data.GclassTree().Register(u)
	return u
}

func (s *User) Name() string {
	return common.REGISTER_MEMORY_USER
}

func (u *User) Register(user *model.User) error {
	gusers.mu.Lock()
	defer gusers.mu.Unlock()
	if _, ok := gusers.list[user.Name]; ok {
		return fmt.Errorf("this user is already exists:%s,", user.Name)
	}
	gusers.list[user.Name] = user
	gusers.locks[user.Name] = &sync.Mutex{}
	return nil
}

func (u *User) Login(name string, pwd string) (model.User, error) {
	muser, ok := gusers.list[name]
	if !ok {
		return emptyUser, fmt.Errorf("this user is not registered:%s，please register", name)
	}
	if muser.Pwd != pwd {
		return emptyUser, errors.New("username or pwd is err")
	}
	gonLineUsers.mu.Lock()
	defer gonLineUsers.mu.Unlock()
	muser.LastLoginTime = tools.NowTimeFormatTimeToString()
	gonLineUsers.list[name] = muser
	cuser := *muser
	cuser.Pwd = ""
	return cuser, nil
}

func (u *User) ForgetPwd(name string, pwd string) error {
	return nil
}

func (u *User) LoginOut(user model.User) error {
	gonLineUsers.mu.Lock()
	defer gonLineUsers.mu.Unlock()
	delete(gonLineUsers.list, user.Name)
	return nil
}

func (u *User) IsLogin(name string) (model.User, bool) {
	gonLineUsers.mu.RLock()
	defer gonLineUsers.mu.RUnlock()
	if model, ok := gonLineUsers.list[name]; ok {
		return *model, true
	}
	return emptyUser, false
}

func (u *User) Update(user model.User) (model.User, error) {

	if err := user.CanUpdate(); err != nil {
		return user, err
	}
	gusers.mu.Lock()
	defer gusers.mu.Unlock()
	_, ok := gusers.list[user.Name]
	if !ok {
		return emptyUser, errors.New("this user is not register")
	}
	gusers.list[user.Name] = &user

	gonLineUsers.mu.Lock()
	defer gonLineUsers.mu.Unlock()
	_, ok = gonLineUsers.list[user.Name]
	if ok {
		gonLineUsers.list[user.Name] = &user
	}

	return user, nil
}

func (u *User) IsOnLine(name string) (model.User, bool) {
	gonLineUsers.mu.RLock()
	defer gonLineUsers.mu.RUnlock()
	user, ok := gonLineUsers.list[name]
	if !ok {
		return emptyUser, false
	}
	return *user, true
}

func (u *User) OnLineUserList() []model.User {
	var res []model.User
	for _, user := range gonLineUsers.list {
		res = append(res, *user)
	}
	return res
}

func (u *User) UserList() []model.User {
	var res []model.User
	for _, user := range gusers.list {
		res = append(res, *user)
	}
	return res
}

func (u *User) FetchUser(name string) (model.User, error) {
	gusers.mu.RLock()
	defer gusers.mu.RUnlock()
	if user, ok := gusers.list[name]; ok {
		return *user, nil
	}
	return emptyUser, errors.New("no this user")
}

func (u *User) UserCanTransformPlayer(name string) (model.User, error) {
	gonLineUsers.mu.RLock()
	defer gonLineUsers.mu.RUnlock()
	user, ok := gonLineUsers.list[name]
	if !ok {
		return emptyUser, errors.New("this user is offline")
	}
	if !user.IsPrepare() {
		return emptyUser, errors.New("this user status  is not prepare")
	}
	return *user, nil
}

func (u *User) Lock(username string) *sync.Mutex {
	return gusers.locks[username]
}
