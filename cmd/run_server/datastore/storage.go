package datastore

import (
	"errors"
	"sync"
)

// Storage represents local runtime storage.
type Storage struct {
	mu sync.RWMutex
	db map[string]bool
}

// NewStorage implements and new storage object.
func NewStorage() *Storage {
	m := make(map[string]bool)
	return &Storage{
		db: m,
	}
}

// AddItem creates a new object in storage.
func (s *Storage) AddItem(in string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.db[in] = true
	if !s.db[in] {
		return errors.New("item not added")
	}
	return nil
}

// GetItem returns a matching object from storage.
func (s *Storage) GetItem(key string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.db[key]
	if !ok || !val {
		return "", errors.New("unable to get item")
	}
	return key, nil
}

// RemoveItem deletes the corresponding object from storage.
func (s *Storage) RemoveItem(key string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	val, ok := s.db[key]
	if !ok || !val {
		return errors.New("cannot remove item that does not exist")
	}

	delete(s.db, key)
	val, ok = s.db[key]
	if ok || val {
		return errors.New("unable to remove item")
	}
	return nil
}

// ListItems returns all objects in storage.
func (s *Storage) ListItems() ([]string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	keys := []string{}
	for k := range s.db {
		keys = append(keys, k)
	}
	if len(keys) < 1 {
		return nil, errors.New("no items to list")
	}
	return keys, nil
}