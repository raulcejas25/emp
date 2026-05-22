package store

import (
	"sync"
	"emp/internal/model"
)

type CategoryStore struct {
	mu         sync.Mutex
	categories []model.Category
	nextID     int
}

func NewCategoryStore() *CategoryStore {
	return &CategoryStore{
		categories: []model.Category{},
		nextID:     1,
	}
}

func (s *CategoryStore) GetAll() []model.Category {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.categories
}

func (s *CategoryStore) Create(name, description string) model.Category {
	s.mu.Lock()
	defer s.mu.Unlock()
	c := model.Category{
		ID:          s.nextID,
		Name:        name,
		Description: description,
	}
	s.nextID++
	s.categories = append(s.categories, c)
	return c
}
