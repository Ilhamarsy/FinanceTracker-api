package services

import (
	"errors"
	"finance-tracker/models"
	"finance-tracker/repositories"
)

type CategoryService struct {
	repo *repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	if err := s.repo.CheckCategory(category); err == nil {
		return errors.New("category available")
	}

	return s.repo.CreateCategory(category)
}

func (s *CategoryService) GetCategories(userID uint) ([]models.Category, error) {
	return s.repo.GetCategories(userID)
}

func (s *CategoryService) DeleteCategory(categoryId, userId uint) error {
	return s.repo.DeleteCategory(categoryId, userId)
}

//FindCategoriesByIDs finds categories by their IDs
