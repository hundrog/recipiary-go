package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DB_CONNECTION_STRING")
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
	err = database.AutoMigrate(
		&Category{},
		&Ingredient{},
		&Recipe{},
		&RecipeIngredients{},
		&Instruction{},
		&Schedule{},
		&Account{},
	)
	if err != nil {
		return
	}

	DB = database
}
