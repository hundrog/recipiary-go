package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=localhost user=developer password=password dbname=recipiary port=5432"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	err = database.SetupJoinTable(&Recipe{}, "Ingredients", &RecipeIngredients{})
	if err != nil {
		return
	}
	err = database.AutoMigrate(&Category{}, &Ingredient{}, &Recipe{}, &RecipeIngredients{})
	if err != nil {
		return
	}

	DB = database
}
