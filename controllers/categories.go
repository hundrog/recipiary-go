package controllers

import (
	"net/http"
	"recipiary/models"

	"github.com/gin-gonic/gin"
)

type CreateCategoryInput struct {
	Name  string `binding:"required"`
	Color string `binding:"required"`
}

type UpdateCategoryInput struct {
	Name  string
	Color string
}

// INDEX
func IndexCategories(c *gin.Context) {
	var categories []models.Category
	models.DB.Where("user_id = ?", CurrentUserID(c)).Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// INDEX
func GetCategory(c *gin.Context) {
	// Get record
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// POST
func CreateCategory(c *gin.Context) {
	// Validate input
	var input CreateCategoryInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Create category
	category := models.Category{Name: input.Name, Color: input.Color}
	models.DB.Create(&category)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// UPDATE
func UpdateCategory(c *gin.Context) {
	// Get record
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	// Validate input
	var input UpdateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update category
	models.DB.Model(&category).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func DeleteCategory(c *gin.Context) {
	// Get record
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	// Delete category
	models.DB.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
