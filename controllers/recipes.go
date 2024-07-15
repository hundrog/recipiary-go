package controllers

import (
	"net/http"
	"recipiary/models"

	"github.com/gin-gonic/gin"
)

type CreateRecipeInput struct {
	Name          string `binding:"required"`
	Description   string `binding:"required"`
	IngredientIds []int  `binding:"required"`
}

type UpdateRecipeInput struct {
	Name        string
	Description string
}

// INDEX
func IndexRecipes(c *gin.Context) {
	var recipes []models.Recipe
	models.DB.Preload("Ingredients.Category").Find(&recipes)

	c.JSON(http.StatusOK, gin.H{"data": recipes})
}

// GET
func GetRecipe(c *gin.Context) {
	//Get Reccord
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

// POST
func CreateRecipe(c *gin.Context) {
	// Validate imput
	var input CreateRecipeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve Ingredients
	var ingredients []models.Ingredient
	if err := models.DB.Find(&ingredients, input.IngredientIds).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "One or more Ingredient doesn't exists"})
		return
	}

	// Create recipe
	recipe := models.Recipe{Name: input.Name, Description: input.Description, Ingredients: ingredients}
	models.DB.Create(&recipe)

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

// UPDATE
func UpdateRecipe(c *gin.Context) {
	//Get Reccord
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Validate imput
	var input UpdateRecipeInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update recipe
	values := models.Recipe{Name: input.Name, Description: input.Description}
	models.DB.Model(&recipe).Updates(&values)

	c.JSON(http.StatusOK, gin.H{"data": recipe})
}

func DeleteRecipe(c *gin.Context) {
	//Get Reccord
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Delete recipe
	models.DB.Delete(&recipe)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
