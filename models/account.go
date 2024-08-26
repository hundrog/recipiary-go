package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserID    string
	UserEmail string
	Provider  string
}
