package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	StartDate time.Time
	FinalDate time.Time
	Recipes   []Recipe `gorm:"many2many:schedule_recipes;"`
	UserID    string
}
