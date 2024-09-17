package services

import (
	"errors"
	"finance-tracker/models"
	"finance-tracker/repositories"
)

type IncomeService struct {
	repo *repositories.IncomeRepository
}

func NewIncomeService(repo *repositories.IncomeRepository) *IncomeService {
	return &IncomeService{repo: repo}
}

func (i *IncomeService) AddIncome(income *models.Income) error {
	return i.repo.AddIncome(income)
}

func (i *IncomeService) GetIncomes(userId uint) ([]models.Income, error) {
	return i.repo.GetIncomes(userId)
}

func (i *IncomeService) UpdateIncome(income *models.Income) error {
	_, err := i.repo.GetIncomeByID(income.ID, income.UserID)
	if err != nil {
		return errors.New("incomes unavailable")
	}

	return i.repo.UpdateIncome(income)
}

func (i *IncomeService) DeleteIncome(incomeId, userId uint) error {
	_, err := i.repo.GetIncomeByID(incomeId, userId)
	if err != nil {
		return errors.New("incomes unavailable")
	}

	return i.repo.DeleteIncome(incomeId, userId)
}

func (i *IncomeService) GetIncomeById(incomeId, userID uint) (*models.Income, error) {
	return i.repo.GetIncomeByID(incomeId, userID)
}
