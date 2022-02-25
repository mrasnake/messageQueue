package main

import (
	"errors"
	"sync"
)

type Storage struct{
	mu sync.RWMutex
	db map[string]bool
}

func NewStorage() *Storage{
	m := make(map[string]bool)
	return &Storage{
		db: m,
	}
}

func (s *Storage) AddItem(in string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.db[in] = true
	if !s.db[in]{
		return errors.New("item not added")
	}
	return nil
}

func (s *Storage) GetItem(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.db[key]
	if !ok || !val{
		return "", errors.New("unable to get item")
	}
	return key, nil
}

func (s *Storage) RemoveItem(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.db[key]
	if !ok || !val{
		return errors.New("cannot remove item that does not exist")
	}

	delete(s.db, key)
	val, ok = s.db[key]
	if ok || val{
		return errors.New("unable to remove item")
	}
	return nil
}

func (s *Storage) ListItems() ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := []string{}
	for k := range s.db {
		keys = append(keys, k)
	}
	if len(keys) < 1{
		return nil, errors.New("no items to list")
	}
	return keys, nil
}