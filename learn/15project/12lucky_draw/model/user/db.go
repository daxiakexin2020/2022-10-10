package user

import (
	"errors"
	"sync"
)

const MAX_USERS = 1000

var (
	gdb *db
)

func init() {
	gdb = &db{
		users: make(map[string]*User),
		mu:    sync.Mutex{},
		cap:   MAX_USERS,
	}
}

type db struct {
	users map[string]*User
	mu    sync.Mutex
	cap   int
}

func (d *db) add(user *User) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.cap == len(d.users) {
		return errors.New("cap is overflow")
	}
	d.users[user.Id] = user
	return nil
}

func (d *db) all() map[string]User {
	res := make(map[string]User, len(d.users))
	for k, v := range d.users {
		res[k] = *v
	}
	return res
}

func (d *db) find(uid string) (User, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	user, ok := d.users[uid]
	if !ok {
		return emptyUser, errors.New("no this user")
	}
	return *user, nil
}

func (d *db) delete(uid string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, ok := d.users[uid]; !ok {
		return errors.New("no this user")
	}
	delete(d.users, uid)
	d.cap -= 1
	return nil
}

func (d *db) update(user *User) (User, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	uid := user.Id
	if _, ok := d.users[uid]; !ok {
		return emptyUser, errors.New("no this user")
	}
	d.users[uid] = user
	return *d.users[uid], nil
}

func (d *db) isOverflow() bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.cap == len(d.users)
}

func (d *db) Count() int {
	d.mu.Lock()
	d.mu.Unlock()
	return len(d.users)
}

func getGDB() *db {
	return gdb
}
