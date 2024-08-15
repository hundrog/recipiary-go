package models

import (
	"gorm.io/gorm"
	"time"
)

type Schedule struct {
	gorm.Model
	StartDate time.Time
	FinalDate time.Time
	Recipes   []Recipe `gorm:"many2many:schedule_recipes;"`
}
