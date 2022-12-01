package _map

import "sync"

type Set struct {
	m    map[int]struct{}
	len  int
	lock sync.Mutex
}

func NewSet(cap int) *Set {
	return &Set{
		m: make(map[int]struct{}, cap),
	}
}

func (s *Set) Add(itme int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m[itme] = struct{}{}
	s.len = len(s.m)
}

func (s *Set) Remove(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.m[item]; ok {
		delete(s.m, item)
		s.len = len(s.m)
	}
}

func (s *Set) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.m = map[int]struct{}{}
	s.len = 0
}

func (s *Set) Has(item int) bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if _, ok := s.m[item]; ok {
		return true
	}
	return false
}

func (s *Set) Len() int {
	return s.len
}

func (s *Set) IsEmpty() bool {
	return len(s.m) == 0
}
