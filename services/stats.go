package services

import (
	"finance-tracker/models"
	"finance-tracker/repositories"
)

type StatService struct {
	expenseRepo *repositories.ExpenseRepository
	incomeRepo  *repositories.IncomeRepository
}

func NewStatService(expenseRepo *repositories.ExpenseRepository, incomeRepo *repositories.IncomeRepository) *StatService {
	return &StatService{expenseRepo, incomeRepo}
}

func (s *StatService) GetStats(userId uint) (*models.Stat, error) {
	// Get total expenses and incomes
	totalExpense, err := s.expenseRepo.GetTotalExpense(userId)
	if err != nil {
		return nil, err
	}

	totalIncome, err := s.incomeRepo.GetTotalIncome(userId)
	if err != nil {
		return nil, err
	}
	// Calculate total balance
	totalBalance := totalIncome - totalExpense

	monthlyExpense, err := s.expenseRepo.GetMonthlyExpense(userId)
	if err != nil {
		return nil, err
	}

	monthlyIncome, err := s.incomeRepo.GetMonthlyIncome(userId)
	if err != nil {
		return nil, err
	}

	monthlyBalance := monthlyIncome - monthlyExpense
	// // Calculate average monthly income and expenses
	// averageMonthlyIncome := totalIncomes / 12
	// averageMonthlyExpenses := totalExpenses / 12

	// // Calculate monthly income growth
	// monthlyIncomeGrowth := averageMonthlyIncome - averageMonthlyExpenses

	// // Calculate annual income growth
	// annualIncomeGrowth := monthlyIncomeGrowth * 12

	// Output stats
	// println("Total Expenses:", totalExpense)
	// println("Total Incomes:", totalIncome)
	// println("Total Balance:", totalBalance)

	// println("Monthly Expenses:", monthlyExpense)
	// println("Monthly Incomes:", monthlyIncome)
	// println("Monthly Balance:", monthlyBalance)

	stat := models.Stat{
		TotalExpense:   totalExpense,
		TotalIncome:    totalIncome,
		TotalBalance:   totalBalance,
		MonthlyExpense: monthlyExpense,
		MonthlyIncome:  monthlyIncome,
		MonthlyBalance: monthlyBalance,
	}

	return &stat, nil
}

func (s *StatService) GetYearlyStats(userID uint, year int) (map[string][]models.MonthlyTotal, error) {
	// Get expense totals by month
	expenseTotals, err := s.expenseRepo.GetMonthlyTotals(userID, year)
	if err != nil {
		return nil, err
	}

	// Get income totals by month
	incomeTotals, err := s.incomeRepo.GetMonthlyTotals(userID, year)
	if err != nil {
		return nil, err
	}

	// Combine both into a single map
	stats := map[string][]models.MonthlyTotal{
		"expenses": expenseTotals,
		"incomes":  incomeTotals,
	}

	return stats, nil
}
