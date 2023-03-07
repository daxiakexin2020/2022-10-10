package memory

import (
	"20red_police/internal/model"
	"errors"
	"fmt"
	"sync"
	"time"
)

type User struct{}

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
	gonce.Do(func() {
		gusers = &users{
			list: map[string]*model.User{},
		}
		gonLineUsers = &onLineUsers{
			list: map[string]*model.User{},
		}
	})
}

func NewUser() *User {
	return &User{}
}

func (u *User) Register(user *model.User) error {
	gusers.mu.Lock()
	defer gusers.mu.Unlock()
	if _, ok := gusers.list[user.Name]; ok {
		return fmt.Errorf("此用户名:%s已经存在", user.Name)
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
	cuser := *muser
	cuser.Pwd = ""
	gonLineUsers.mu.Lock()
	defer gonLineUsers.mu.Unlock()
	muser.LastLoginTime = time.Now().Format("2006-01-02 15:04:05")
	gonLineUsers.list[name] = muser
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

func (u *User) Update(user model.User) error {
	return nil
}

func (u *User) isOnLine(name string) bool {
	gonLineUsers.mu.RLock()
	defer gonLineUsers.mu.RUnlock()
	_, ok := gonLineUsers.list[name]
	return ok
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
