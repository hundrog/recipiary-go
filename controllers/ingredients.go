package controllers

import (
	"net/http"
	"recipiary/models"

	"github.com/gin-gonic/gin"
)

type CreateIngredientInput struct {
	Name       string `binding:"required"`
	Portion    string `binding:"required"`
	CategoryId int    `binding:"required"`
}

type UpdateIngredientInput struct {
	Name       string
	Portion    string
	CategoryId int
}

// INDEX
func IndexIngredients(c *gin.Context) {
	var ingredients []models.Ingredient
	models.DB.Joins("Category").Find(&ingredients)

	c.JSON(http.StatusOK, gin.H{"data": ingredients})
}

// GET
func GetIngredient(c *gin.Context) {
	// Get Record
	var ingredient models.Ingredient
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ingredient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ingredient})
}

// POST
func CreateIngredient(c *gin.Context) {
	// Validate input
	var input CreateIngredientInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve category
	var category models.Category
	if err := models.DB.Where("id = ?", input.CategoryId).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category doesn't exists"})
		return
	}

	// Create ingredient
	ingredient := models.Ingredient{Name: input.Name, CategoryId: input.CategoryId, Portion: input.Portion, Category: category}
	models.DB.Create(&ingredient)

	c.JSON(http.StatusOK, gin.H{"data": ingredient})
}

// UPDATE
func UpdateIngredient(c *gin.Context) {
	// Get Record
	var ingredient models.Ingredient
	if err := models.DB.First(&ingredient, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Validate input
	var input UpdateIngredientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve category
	categoryId := input.CategoryId
	if input.CategoryId <= 0 {
		categoryId = ingredient.CategoryId
	}

	var category models.Category
	if err := models.DB.Where("id = ?", categoryId).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category doesn't exists"})
		return
	}

	// Update ingredient
	err := models.DB.Model(&ingredient).Updates(input)
	if err != nil {
		ingredient.Category = category
	}

	c.JSON(http.StatusOK, gin.H{"data": ingredient})
}

func DeleteIngredient(c *gin.Context) {
	// Get Record
	var ingredient models.Ingredient
	if err := models.DB.Where("id = ?", c.Param("id")).First(&ingredient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Delete ingredient
	models.DB.Delete(&ingredient)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
