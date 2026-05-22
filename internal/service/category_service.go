package service

import (
	"errors"
	"emp/internal/model"
	"emp/internal/store"
)

type CategoryService struct {
	store *store.CategoryStore
}

func NewCategoryService(s *store.CategoryStore) *CategoryService {
	return &CategoryService{store: s}
}
func (s *CategoryService) GetCategories() []model.Category {
	return s.store.GetAll()
}
func (s *CategoryService) CreateCategory(name, description string) (model.Category, error) {
	if name == "" {
		return model.Category{}, errors.New("name is required")
	}
	return s.store.Create(name, description), nil
}
