package models

import "gorm.io/gorm"

type Ingredient struct {
	gorm.Model
	Name       string
	Portion    string
	CategoryId int
	Category   Category
	UserID     string
}
