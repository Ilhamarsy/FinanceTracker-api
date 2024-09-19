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
	// 1. Retrieve the existing expense
	existingExpense, err := e.expenseRepository.GetExpenseByID(req.Id, req.UserID)
	if err != nil {
		return errors.New("expense unavailable")
	}

	// 2. Retrieve the current categories associated with the expense
	err = e.categoryRepository.PreloadCategories(existingExpense)
	if err != nil {
		return err
	}
	currentCategories := existingExpense.Categories

	// 3. Find the new categories from the request
	newCategories, err := e.categoryRepository.FindCategoriesByIDs(req.UserID, req.CategoryIDs)
	if err != nil {
		return err
	}

	// 4. Find categories to remove (those in currentCategories but not in newCategories)
	var categoriesToRemove []models.Category
	for _, currentCat := range currentCategories {
		found := false
		for _, newCat := range newCategories {
			if currentCat.ID == newCat.ID {
				found = true
				break
			}
		}
		if !found {
			categoriesToRemove = append(categoriesToRemove, currentCat)
		}
	}

	// 5. Remove the unwanted categories from the expense
	if len(categoriesToRemove) > 0 {
		var expense models.Expense
		expense.ID = req.Id
		err := e.categoryRepository.RemoveCategoriesFromExpense(&expense, categoriesToRemove)
		if err != nil {
			return err
		}
	}

	// 6. Update the expense with new data
	expense := models.Expense{
		Title:       req.Title,
		Description: req.Description,
		Amount:      req.Amount,
		UserID:      req.UserID,
		Categories:  newCategories,
	}

	if len(req.CategoryIDs) < 1 {
		var emptyCategory []models.Category
		expense.Categories = emptyCategory
	}

	expense.ID = req.Id

	// 7. Update the expense record
	err = e.expenseRepository.UpdateExpense(&expense)
	if err != nil {
		return err
	}

	return nil
}

// func (e *ExpenseService) UpdateExpense(req *models.ExpenseRequest) error {
// 	_, err := e.expenseRepository.GetExpenseByID(req.Id, req.UserID)
// 	if err != nil {
// 		return errors.New("expense unavailable")
// 	}

// 	expense := models.Expense{
// 		Title:       req.Title,
// 		Description: req.Description,
// 		Amount:      req.Amount,
// 		UserID:      req.UserID,
// 	}

// 	categories, err := e.categoryRepository.FindCategoriesByIDs(req.UserID, req.CategoryIDs)
// 	if err != nil {
// 		return err
// 	}

// 	fmt.Println(categories)

// 	expense.ID = req.Id
// 	expense.Categories = categories

// 	return e.expenseRepository.UpdateExpense(&expense)
// }

func (e *ExpenseService) DeleteExpense(expenseID, userID uint) error {
	_, err := e.expenseRepository.GetExpenseByID(expenseID, userID)
	if err != nil {
		return errors.New("expense unavailable")
	}

	return e.expenseRepository.DeleteExpense(expenseID, userID)
}
