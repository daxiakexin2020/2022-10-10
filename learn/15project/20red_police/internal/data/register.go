package data

import (
	"errors"
	"sync"
)

type classTree struct {
	list map[string]Class
	mu   sync.RWMutex
}

var gclassTree *classTree

func init() {
	gclassTree = &classTree{list: map[string]Class{}}
}

func (cr *classTree) Register(cs ...Class) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	for _, c := range cs {
		if _, ok := cr.list[c.Name()]; ok {
			return errors.New("Class Name is Registered:" + c.Name())
		}
		cr.list[c.Name()] = c
	}
	return nil
}

func (cr *classTree) Pick(cname string) (Class, error) {
	cr.mu.RLock()
	defer cr.mu.RUnlock()
	if class, ok := cr.list[cname]; ok {
		return class, nil
	}
	return nil, errors.New("this class is not registered:" + cname)
}

func GclassTree() *classTree {
	return gclassTree
}

//func FileUser() (*stores.User, error) {
//	pick, err := GclassTree().Pick(common.REGISTER_FILE_DATA_USER)
//	if err != nil {
//		return nil, err
//	}
//	user, ok := pick.(*stores.User)
//	if !ok {
//		return nil, errors.New("file user class is err")
//	}
//	return user, nil
//}

//func MemoryUser() (*memory.User, error) {
//	pick, err := GclassTree().Pick(common.REGISTER_DATA_USER)
//	if err != nil {
//		return nil, err
//	}
//	user, ok := pick.(*memory.User)
//	if !ok {
//		return nil, errors.New("file user class is err")
//	}
//	return user, nil
//}

//func MemoryRoom() (*memory.Room, error) {
//	pick, err := GclassTree().Pick(common.REGISTER_DATA_ROOM)
//	if err != nil {
//		return nil, err
//	}
//	room, ok := pick.(*memory.Room)
//	if !ok {
//		return nil, errors.New("memory room class is err")
//	}
//	return room, nil
//}
