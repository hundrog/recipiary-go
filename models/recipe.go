package models

import "gorm.io/gorm"

type Recipe struct {
	gorm.Model
	Name         string
	Description  string
	ImageUrl     string
	Ingredients  []Ingredient `gorm:"many2many:recipe_ingredients;"`
	Instructions []Instruction
}

type RecipeIngredients struct {
	RecipeID     int `gorm:"primaryKey"`
	IngredientID int `gorm:"primaryKey"`
	Amount       int `default:"0"`
}

type Instruction struct {
	gorm.Model
	Content  string
	RecipeID int
}
