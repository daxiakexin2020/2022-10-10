package memory

import (
	"20red_police/internal/data"
	"20red_police/internal/model"
	"20red_police/tools"
	"errors"
	"fmt"
	"sync"
)

type User struct{}

var _ data.User = (*User)(nil)

type users struct {
	list map[string]*model.User
	mu   sync.RWMutex
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
			list: make(map[string]*model.User),
		}
		gonLineUsers = &onLineUsers{
			list: make(map[string]*model.User),
		}
	})
}

func NewUser() data.User {
	return &User{}
}

func (s *User) Name() string {
	return "User"
}

func (u *User) Register(user *model.User) error {
	gusers.mu.Lock()
	defer gusers.mu.Unlock()
	if _, ok := gusers.list[user.Name]; ok {
		return fmt.Errorf("此用户名:%s,已经存在", user.Name)
	}
	gusers.list[user.Name] = user
	return nil
}

func (u *User) Login(name string, pwd string) (model.User, error) {
	muser, ok := gusers.list[name]
	if !ok {
		return emptyUser, fmt.Errorf("此用户:%s没有注册，请先注册", name)
	}
	if muser.Pwd != pwd {
		return emptyUser, errors.New("用户名或密码错误")
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

func (u *User) Update(user model.User) error {
	return nil
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
