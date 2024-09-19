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

	if totalIncome, ok := income["total_income"]; ok && totalIncome != nil {
		return totalIncome.(float64), nil
	}

	return 0, nil
}

func (i *IncomeRepository) GetTotalIncome(userID uint) (float64, error) {
	var income map[string]interface{}
	query := `SELECT SUM(amount) AS total_income 
						FROM incomes 
						WHERE user_id = ? 
						AND deleted_at IS NULL`

	err := i.DB.Raw(query, userID).Scan(&income).Error
	if err != nil {
		return 0, err
	}

	if totalIncome, ok := income["total_income"]; ok && totalIncome != nil {
		return totalIncome.(float64), nil
	}

	return 0, nil
}

func (e *IncomeRepository) GetMonthlyTotals(userID uint, year int) ([]models.MonthlyTotal, error) {
	var monthlyTotals []models.MonthlyTotal

	err := e.DB.
		Table("incomes").
		Select("EXTRACT(MONTH FROM created_at) as month, SUM(amount) as total").
		Where("user_id = ? AND EXTRACT(YEAR FROM created_at) = ?", userID, year).
		Group("EXTRACT(MONTH FROM created_at)").
		Order("EXTRACT(MONTH FROM created_at)").
		Scan(&monthlyTotals).Error

	if err != nil {
		return nil, err
	}

	return monthlyTotals, nil
}
