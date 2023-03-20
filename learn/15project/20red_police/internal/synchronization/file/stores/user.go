package stores

import (
	"20red_police/common"
	"20red_police/internal/data"
	"20red_police/internal/data/memory"
	"20red_police/internal/model"
	"20red_police/internal/synchronization/file"
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
	file, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	return &User{store: store{file: file}}, nil
}

func (u *User) Read() error {
	//u.mu.Lock()
	//defer u.mu.Unlock()
	//pick, err := data.MemoryUser()
	//if err != nil {
	//	return err
	//}
	//
	//muser, ok := pick.(*memory.User)
	//if !ok {
	//	return errors.New("memory class err:")
	//}

	//var err error
	//for err != io.EOF {
	//	 u.store.read()
	//	if err != nil {
	//		log.Println("file write user err:", err)
	//	} else {
	//		log.Println("write ok:", n)
	//	}
	//}
	return nil

}

func (u *User) Write(wdata interface{}) error {
	if wdata == nil {
		return u.bacthWrite()
	}
	return u.write(wdata)
}

func (u *User) write(wdata interface{}) error {
	u.mu.Lock()
	defer u.mu.Unlock()
	_, err := u.store.write(wdata)
	return err
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

	list := muser.UserList()
	for _, line := range list {
		fmt.Println("line:", line)
		n, err := u.store.write(line)
		if err != nil {
			log.Println("file write user err:", err)
		} else {
			log.Println("write ok:", n)
		}
	}
	return nil
}

func (u *User) Close() error {
	return u.store.close()
}

func (u *User) Name() string {
	return common.REGISTER_FILE_DATA_USER
}

func MemoryBatchSyncFile() error {
	pick, err := data.FileUser()
	if err != nil {
		return err
	}
	user, ok := pick.(*User)
	if !ok {
		return errors.New("file user class is err")
	}
	return user.Write(nil)
}

func MemorySyncFile(model *model.User) error {
	pick, err := data.FileUser()
	if err != nil {
		return err
	}
	user, ok := pick.(*User)
	if !ok {
		return errors.New("file user class is err")
	}
	return user.Write(model)
}

func FileSyncMemory() error {
	pick, err := data.FileUser()
	if err != nil {
		return err
	}
	user, ok := pick.(*User)
	if !ok {
		return errors.New("file user class is err")
	}
	return user.Read()
}
