package stores

import (
	"20red_police/common"
	"20red_police/internal/data"
	"20red_police/internal/data/memory"
	"20red_police/internal/model"
	"20red_police/internal/synchronization/file"
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"sync"
)

type User struct {
	store
	mu sync.Mutex
}

var _ file.WR = (*User)(nil)

const user_file_path = "user.txt"

func NewUser() (*User, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path := pwd + "/internal/synchronization/file/txt/" + user_file_path
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		return nil, err
	}
	return &User{store: store{file: file, reader: bufio.NewReader(file), writer: bufio.NewWriter(file)}}, nil
}

func (u *User) Read() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	pick, err := data.MemoryUser()
	if err != nil {
		return err
	}
	muser, ok := pick.(*memory.User)
	if !ok {
		return errors.New("memory class err:")
	}

	var users []*model.User
	err = u.store.read(func(buf []byte) {
		var user model.User
		if err = json.Unmarshal(buf, &user); err != nil {
			fmt.Println("read json unmarshal err:", err)
		} else {
			users = append(users, &user)
		}
	})
	if err != nil {
		return err
	}
	for _, user := range users {
		if err = muser.Register(user); err != nil {
			fmt.Println("file into memory register user err:", err)
		}
		muser.Login(user.Name, user.Pwd)
	}
	log.Println("File user data read into memory ok......................")
	return nil
}

func (u *User) Write(wdata interface{}) error {
	if wdata == nil {
		return u.bacthWrite()
	}
	return u.write(wdata)
}

func (u *User) write(data interface{}) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	line, err := json.Marshal(data)
	if err != nil {
		log.Println("file marshal err:", err)
		return err
	}
	line = append(line, '\n')
	u.store.write(line)
	return u.store.flush()
}

func (u *User) bacthWrite() error {
	u.mu.Lock()
	defer u.mu.Unlock()
	pick, err := data.MemoryUser()
	if err != nil {
		return err
	}
	muser, ok := pick.(*memory.User)
	if !ok {
		return errors.New("memory class err:")
	}

	var lines []byte
	list := muser.UserList()

	for _, user := range list {
		line, err := json.Marshal(user)
		if err != nil {
			log.Println("file marshal err:", err)
			continue
		}
		line = append(line, '\n')
		lines = append(lines, line...)
	}
	u.store.write(lines)
	return u.store.flush()
}

func (u *User) Close() error {
	return u.store.close()
}

func (u *User) Name() string {
	return common.REGISTER_FILE_USER
}

func UserMemoryBatchSyncFile() error {
	pick, err := data.FileUser()
	if err != nil {
		return err
	}
	user, ok := pick.(*User)
	if !ok {
		return errors.New("UserMemoryBatchSyncFile file user class is err")
	}
	return user.Write(nil)
}

func UserMemorySyncFile(model interface{}) error {
	pick, err := data.FileUser()
	if err != nil {
		return err
	}
	user, ok := pick.(*User)
	if !ok {
		return errors.New("UserMemorySyncFile file user class is err")
	}
	return user.Write(model)
}

func CloseUserSyncFile() error {
	pick, err := data.FileUser()
	if err != nil {
		return err
	}
	user, ok := pick.(*User)
	if !ok {
		return errors.New("CloseUserSyncFile file user class is err")
	}
	return user.store.close()
}

func UserFileSyncMemory() error {
	pick, err := data.FileUser()
	if err != nil {
		return err
	}
	user, ok := pick.(*User)
	if !ok {
		return errors.New("UserFileSyncMemory file user class is err")
	}

	return user.Read()
}
