package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recipiary/models"
)

type CreateCategoryImput struct {
	Name  string `binding:"required"`
	Color string `binding:"required"`
}

type UpdateCategoryImput struct {
	Name  string
	Color string
}

// INDEX
func IndexCategories(c *gin.Context) {
	var categories []models.Category
	models.DB.Find(&categories)

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// GET
func GetCategory(c *gin.Context) {
	//Get Reccord
	var category models.Category
	if err := models.DB.First(&category, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// POST
func CreateCategory(c *gin.Context) {
	// Validate imput
	var input CreateCategoryImput

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
	//Get Reccord
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Validate imput
	var input UpdateCategoryImput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update category
	models.DB.Model(&category).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": category})
}

func DeleteCategory(c *gin.Context) {
	//Get Reccord
	var category models.Category
	if err := models.DB.Where("id = ?", c.Param("id")).First(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Delete category
	models.DB.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
