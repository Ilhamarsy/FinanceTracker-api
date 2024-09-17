package models

type Stat struct {
	TotalBalance   float64 `json:"total_balance"`
	TotalIncome    float64 `json:"total_income"`
	TotalExpense   float64 `json:"total_expense"`
	MonthlyBalance float64 `json:"monthly_balance"`
	MonthlyIncome  float64 `json:"monthly_income"`
	MonthlyExpense float64 `json:"monthly_expense"`
	// RecentIncome  Income  `json:"recent_income"`
	// RecentExpense Expense `json:"recent_expense"`
}
