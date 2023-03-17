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

func (cr *classTree) Register(c Class) error {
	cr.mu.Lock()
	defer cr.mu.Unlock()
	if _, ok := cr.list[c.Name()]; ok {
		return errors.New("Class Name is Registered:" + c.Name())
	}
	cr.list[c.Name()] = c
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
