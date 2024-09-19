package repositories

import (
	"finance-tracker/models"
	"fmt"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	DB *gorm.DB
}

func NewExpenseRepository(db *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{DB: db}
}

func (e *ExpenseRepository) AddExpense(expense *models.Expense) error {
	return e.DB.Create(expense).Error
}

func (e *ExpenseRepository) GetExpenses(userID uint) ([]models.Expense, error) {
	var expenses []models.Expense
	return expenses, e.DB.Preload("Categories").Where("user_id = ?", userID).Find(&expenses).Error
}

func (e *ExpenseRepository) UpdateExpense(expense *models.Expense) error {
	return e.DB.Model(expense).Omit("CreatedAt").Updates(expense).Error
}

func (e *ExpenseRepository) DeleteExpense(id, userId uint) error {
	return e.DB.Where("id = ? AND user_id = ?", id, userId).Delete(&models.Expense{}).Error
}

func (e *ExpenseRepository) GetExpenseByID(id, userId uint) (*models.Expense, error) {
	var expense models.Expense
	err := e.DB.Where("id = ? AND user_id = ?", id, userId).First(&expense).Error
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func (e *ExpenseRepository) GetMonthlyExpense(userID uint) (float64, error) {
	var expense map[string]interface{}
	query := `SELECT EXTRACT(MONTH FROM created_at) AS month, EXTRACT(YEAR FROM created_at) AS year, SUM(amount) AS total_expense 
						FROM expenses 
						WHERE user_id = ? 
						AND deleted_at IS NULL 
						AND EXTRACT(MONTH FROM created_at) = EXTRACT(MONTH FROM CURRENT_DATE) 
						AND EXTRACT(YEAR FROM created_at) = EXTRACT(YEAR FROM CURRENT_DATE) 
						GROUP BY EXTRACT(MONTH FROM created_at), EXTRACT(YEAR FROM created_at)`

	err := e.DB.Raw(query, userID).Scan(&expense).Error
	if err != nil {
		return 0, err
	}

	fmt.Println(expense["total_expense"])

	if totalExpense, ok := expense["total_expense"]; ok && totalExpense != nil {
		return totalExpense.(float64), nil
	}

	return 0, nil
}

func (e *ExpenseRepository) GetTotalExpense(userID uint) (float64, error) {
	var expense map[string]interface{}
	query := `SELECT SUM(amount) as total_expense
						FROM expenses 
						WHERE user_id = ? 
						AND deleted_at IS NULL`

	err := e.DB.Raw(query, userID).Scan(&expense).Error
	if err != nil {
		return 0, err
	}

	if totalExpense, ok := expense["total_expense"]; ok && totalExpense != nil {
		return totalExpense.(float64), nil
	}

	return 0, nil
}

func (e *ExpenseRepository) GetMonthlyTotals(userID uint, year int) ([]models.MonthlyTotal, error) {
	var monthlyTotals []models.MonthlyTotal

	err := e.DB.
		Table("expenses").
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
