package models

import "gorm.io/gorm"

type Income struct {
	gorm.Model
	Title       string  `json:"title"`
	Description string  `json:"desc"`
	Amount      float64 `json:"amount"`
	User        User    `json:"-"`
	UserID      uint    `json:"-"`
}
