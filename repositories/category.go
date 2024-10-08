package repositories

import (
	"finance-tracker/models"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoryRepository) GetCategories(userID uint) ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Where("user_id = ?", userID).Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) CheckCategory(category *models.Category) error {
	return r.DB.Where("name = ? AND user_id = ?", category.Name, category.UserID).First(category).Error
}

func (r *CategoryRepository) FindCategoriesByIDs(userId uint, ids []uint) ([]models.Category, error) {
	var categories []models.Category
	if err := r.DB.Where("user_id = ?", userId).Find(&categories, ids).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (r *CategoryRepository) DeleteCategory(id, userId uint) error {
	return r.DB.Where("id =? AND user_id =?", id, userId).Delete(&models.Category{}).Error
}

func (r *CategoryRepository) PreloadCategories(expense *models.Expense) error {
	return r.DB.Preload("Categories").Find(expense).Error
}

func (r *CategoryRepository) RemoveCategoriesFromExpense(expense *models.Expense, categories []models.Category) error {
	return r.DB.Model(expense).Association("Categories").Delete(categories)
}
