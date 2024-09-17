package models

import "gorm.io/gorm"

type Expense struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"desc"`
	Amount      float64    `json:"amount"`
	Categories  []Category `gorm:"many2many:expense_categories;" json:"categories"`
	User        User       `json:"-"`
	UserID      uint       `json:"-"`
}

type ExpenseRequest struct {
	Id          uint
	Title       string  `json:"title"`
	Description string  `json:"desc"`
	Amount      float64 `json:"amount"`
	CategoryIDs []uint  `json:"category_ids"`
	UserID      uint
}
