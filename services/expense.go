package services

import (
	"errors"
	"finance-tracker/models"
	"finance-tracker/repositories"
)

type ExpenseService struct {
	expenseRepository  *repositories.ExpenseRepository
	categoryRepository *repositories.CategoryRepository
}

func NewExpenseService(expenseRepository *repositories.ExpenseRepository, categoryRepository *repositories.CategoryRepository) *ExpenseService {
	return &ExpenseService{expenseRepository: expenseRepository, categoryRepository: categoryRepository}
}

func (e *ExpenseService) AddExpense(req models.ExpenseRequest) error {
	expense := models.Expense{
		Title:       req.Title,
		Description: req.Description,
		Amount:      req.Amount,
		UserID:      req.UserID,
	}

	categories, err := e.categoryRepository.FindCategoriesByIDs(req.UserID, req.CategoryIDs)
	if err != nil {
		return err
	}

	expense.Categories = categories

	return e.expenseRepository.AddExpense(&expense)
}

func (e *ExpenseService) GetExpenses(userID uint) ([]models.Expense, error) {
	return e.expenseRepository.GetExpenses(userID)
}

func (e *ExpenseService) UpdateExpense(req *models.ExpenseRequest) error {
	_, err := e.expenseRepository.GetExpenseByID(req.Id, req.UserID)
	if err != nil {
		return errors.New("expense unavailable")
	}

	expense := models.Expense{
		Title:       req.Title,
		Description: req.Description,
		Amount:      req.Amount,
		UserID:      req.UserID,
	}

	categories, err := e.categoryRepository.FindCategoriesByIDs(req.UserID, req.CategoryIDs)
	if err != nil {
		return err
	}

	expense.ID = req.Id
	expense.Categories = categories

	return e.expenseRepository.UpdateExpense(&expense)
}

func (e *ExpenseService) DeleteExpense(expenseID, userID uint) error {
	_, err := e.expenseRepository.GetExpenseByID(expenseID, userID)
	if err != nil {
		return errors.New("expense unavailable")
	}

	return e.expenseRepository.DeleteExpense(expenseID, userID)
}
