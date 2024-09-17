package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	User     User      `json:"-"`
	UserID   uint      `json:"-"`
	Expenses []Expense `gorm:"many2many:expense_categories;" json:"expenses"`
}
