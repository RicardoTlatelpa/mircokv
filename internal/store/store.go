package store

import (
	"sync"
)

type Store struct {
	data map[string]string
	mu sync.RWMutex
}

func NewStore() *Store {
	return &Store {
		data: make(map[string]string),
	}
}

func (s *Store) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]

	return val,ok
}

func (s *Store) Set(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}
