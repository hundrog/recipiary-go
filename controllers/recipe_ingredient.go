package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recipiary/models"
	"strconv"
)

type InsertRecipeIngredientInput struct {
	IngredientId int `binding:"required"`
	Amount       int `binding:"required"`
}

type UpdateRecipeIngredientInput struct {
	Amount int `binding:"required"`
}

// POST
func InsertRecipeIngredient(c *gin.Context) {
	// Validate Input
	var input InsertRecipeIngredientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//Get Reccords
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe not found"})
		return
	}
	var ingredient models.Ingredient
	if err := models.DB.First(&ingredient, input.IngredientId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingredient not found"})
		return
	}
	// Insert recipe ingredient
	recipeIngredient := models.RecipeIngredients{RecipeID: int(recipe.ID), IngredientID: int(ingredient.ID), Amount: input.Amount}

	models.DB.Create(&recipeIngredient)
	c.JSON(http.StatusOK, gin.H{"data": recipeIngredient})
}

func UpdateRecipeIngredient(c *gin.Context) {
	// Validate IDs
	recipeID, err1 := strconv.Atoi(c.Param("id"))
	ingredientID, err2 := strconv.Atoi(c.Param("ingredientID"))
	if err2 != nil || err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe or Ingredient is not valid"})
		return
	}
	// Validate Input
	var input UpdateRecipeIngredientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Fetch recipe ingredient
	var recipeIngredient models.RecipeIngredients
	if err := models.DB.Where(&models.RecipeIngredients{
		RecipeID:     recipeID,
		IngredientID: ingredientID,
	}).First(&recipeIngredient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe or Ingredient not found"})
		return
	}
	// Update recipe ingredient
	models.DB.Model(&recipeIngredient).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": recipeIngredient})
}

// DELETE
func DeleteRecipeIngredient(c *gin.Context) {
	//Get Reccords
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe not found"})
		return
	}
	var ingredient models.Ingredient
	if err := models.DB.First(&ingredient, c.Param("ingredient_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingredient not found"})
		return
	}
	// Delete recipe ingredient
	models.DB.Model(&recipe).Association("Ingredients").Delete(ingredient)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
