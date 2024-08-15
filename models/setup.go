package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
	"log"
	"os"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
	err = database.AutoMigrate(&Category{}, &Ingredient{}, &Recipe{}, &RecipeIngredients{}, &Instruction{}, &Schedule{})
	if err != nil {
		return
	}

	DB = database
}
