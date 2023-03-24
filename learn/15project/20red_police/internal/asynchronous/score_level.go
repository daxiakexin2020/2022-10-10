package asynchronous

import (
	"20red_police/internal/data"
	"20red_police/internal/data/memory"
	"errors"
	"fmt"
	"sync"
)

func HandleScoreLevel(username string, score int64) error {
	if err := upgrade(username, score); err != nil {
		fmt.Println("HandleScoreLevel err:", err)
	}
	return nil
}

var m sync.Mutex

func upgrade(username string, score int64) error {
	//m.Lock()
	//defer m.Unlock()
	pick, err := data.MemoryUser()
	if err != nil {
		return err
	}
	user, ok := pick.(*memory.User)
	if !ok {
		return errors.New("memory class is err")
	}

	fetchUser, err := user.FetchUser(username)
	if err != nil {
		return err
	}
	if err = fetchUser.Upgrade(score); err != nil {
		fmt.Println("user Upgrade err:", err)
	}

	_, err = user.Update(fetchUser)
	if err != nil {
		return err
	}
	return nil
}
