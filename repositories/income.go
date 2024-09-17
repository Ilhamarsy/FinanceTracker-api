package repositories

import (
	"finance-tracker/models"

	"gorm.io/gorm"
)

type IncomeRepository struct {
	DB *gorm.DB
}

func NewIncomeRepository(db *gorm.DB) *IncomeRepository {
	return &IncomeRepository{DB: db}
}

func (i *IncomeRepository) AddIncome(income *models.Income) error {
	return i.DB.Create(income).Error
}

func (i *IncomeRepository) GetIncomes(userID uint) ([]models.Income, error) {
	var incomes []models.Income
	return incomes, i.DB.Where("user_id = ?", userID).Find(&incomes).Error
}

func (i *IncomeRepository) UpdateIncome(income *models.Income) error {
	return i.DB.Model(income).Omit("CreatedAt").Updates(income).Error
}

func (i *IncomeRepository) DeleteIncome(id, userId uint) error {
	return i.DB.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Income{}).Error
}

func (i *IncomeRepository) GetIncomeByID(id, userId uint) (*models.Income, error) {
	var income models.Income
	err := i.DB.Where("id = ? AND user_id = ?", id, userId).First(&income).Error
	if err != nil {
		return nil, err
	}
	return &income, nil
}

func (i *IncomeRepository) GetMonthlyIncome(userID uint) (float64, error) {
	var income map[string]interface{}

	query := `SELECT EXTRACT(MONTH FROM created_at) AS month, EXTRACT(YEAR FROM created_at) AS year, SUM(amount) AS total_income 
						FROM incomes 
						WHERE user_id = ? 
						AND deleted_at IS NULL 
						AND EXTRACT(MONTH FROM created_at) = EXTRACT(MONTH FROM CURRENT_DATE) 
						AND EXTRACT(YEAR FROM created_at) = EXTRACT(YEAR FROM CURRENT_DATE) 
						GROUP BY EXTRACT(MONTH FROM created_at), EXTRACT(YEAR FROM created_at)`

	err := i.DB.Raw(query, userID).Scan(&income).Error
	if err != nil {
		return 0, err
	}

	return income["total_income"].(float64), nil
}

func (i *IncomeRepository) GetTotalIncome(userID uint) (float64, error) {
	var total_amount float64
	query := `SELECT SUM(amount) AS total_income 
						FROM incomes 
						WHERE user_id = ? 
						AND deleted_at IS NULL`

	err := i.DB.Raw(query, userID).Scan(&total_amount).Error
	if err != nil {
		return 0, err
	}

	return total_amount, nil
}
