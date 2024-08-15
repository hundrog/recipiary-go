package controllers

import (
	"fmt"
	"net/http"
	"recipiary/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type ScheduleRecipesInput struct {
	RecipeIDs []int `binding:"required,min=1"`
}

// POST
func CreateScheduleRecipes(c *gin.Context) {
	// Get records
	schedule, recipes, err := getScheduleAndRecipes(c)
	if err != nil {
		handleError(c, err)
		return
	}
	// Update Associations
	models.DB.Model(&schedule).Association("Recipes").Append(recipes)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// PATCH
func UpdateScheduleRecipes(c *gin.Context) {
	// Get records
	schedule, recipes, err := getScheduleAndRecipes(c)
	if err != nil {
		handleError(c, err)
		return
	}
	// Update Associations
	models.DB.Model(&schedule).Association("Recipes").Replace(recipes)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// DELETE
func DeleteScheduleRecipes(c *gin.Context) {
	// Get records
	schedule, recipes, err := getScheduleAndRecipes(c)
	if err != nil {
		handleError(c, err)
		return
	}
	// Update Associations
	models.DB.Model(&schedule).Association("Recipes").Delete(recipes)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// DELETE
func ClearScheduleRecipes(c *gin.Context) {
	// Get records
	var schedule models.Schedule
	if err := models.DB.First(&schedule, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update Associations
	models.DB.Model(&schedule).Association("Recipes").Clear()

	c.JSON(http.StatusOK, gin.H{"data": true})
}

func getScheduleAndRecipes(c *gin.Context) (schedule models.Schedule, recipes []models.Recipe, error error) {
	// Validate input
	var input ScheduleRecipesInput
	if err := c.ShouldBindJSON(&input); err != nil {
		error = fmt.Errorf("invalid input: %w", err)
		return
	}
	// Get Schedule
	if err := models.DB.First(&schedule, c.Param("id")).Error; err != nil {
		error = fmt.Errorf("schedule not found: %w", err)
		return
	}

	// Get Recipes
	if err := models.DB.Find(&recipes, input.RecipeIDs).Error; err != nil {
		return schedule, recipes, fmt.Errorf("error fetching recipes: %w", err)
	}

	return
}
func handleError(c *gin.Context, err error) {
	error := err.Error()
	switch {
	case strings.Contains(error, "schedule not found"), strings.Contains(error, "some recipes were not found"):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	case strings.Contains(error, "invalid input"):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	}
}
