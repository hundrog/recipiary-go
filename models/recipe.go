package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name        string
	Description string
	Ingredients []Ingredient `gorm:"many2many:recipe_ingredients;"`
}

type RecipeIngredients struct {
	RecipeID     int `gorm:"primaryKey"`
	IngredientID int `gorm:"primaryKey"`
	Amount       int `default:"0"`
}
