package controllers

import (
	"net/http"
	"recipiary/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InsertRecipeIngredientInput struct {
	ID     int `binding:"required"`
	Amount int `binding:"required"`
}

type UpdateRecipeIngredientInput struct {
	Amount int `binding:"required"`
}

type RecipeIngredientsResponse struct {
	ID      int
	Name    string
	Portion string
	Amount  int
}

// GET
func GetRecipeIngredients(c *gin.Context) {
	var recipeIngredients []RecipeIngredientsResponse

	models.DB.Raw(`
        SELECT
		ingredients.id,
		ingredients.name AS name,
		ingredients.portion AS portion,
		recipe_ingredients.recipe_id,
		recipe_ingredients.ingredient_id,
		recipe_ingredients.amount AS amount
        FROM 
            recipe_ingredients
        JOIN 
            ingredients ON recipe_ingredients.ingredient_id = ingredients.id
        WHERE 
            recipe_ingredients.recipe_id = ?`, c.Param("id")).Scan(&recipeIngredients)

	c.JSON(http.StatusOK, gin.H{"data": recipeIngredients})
}

// POST
func InsertRecipeIngredient(c *gin.Context) {
	// Validate Input
	var input InsertRecipeIngredientInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Get records
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe not found"})
		return
	}
	var ingredient models.Ingredient
	if err := models.DB.First(&ingredient, input.ID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingredient not found"})
		return
	}
	// Insert recipe ingredient
	recipeIngredient := models.RecipeIngredients{RecipeID: int(recipe.ID), IngredientID: int(ingredient.ID), Amount: input.Amount}
	if err := models.DB.Create(&recipeIngredient).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	responseIngredient := RecipeIngredientsResponse{ID: int(ingredient.ID), Name: ingredient.Name, Portion: ingredient.Portion, Amount: recipeIngredient.Amount}
	c.JSON(http.StatusOK, gin.H{"data": responseIngredient})
}

func UpdateRecipeIngredient(c *gin.Context) {
	// Validate IDs
	recipeID, err1 := strconv.Atoi(c.Param("id"))
	ingredientID, err2 := strconv.Atoi(c.Param("ingredientId"))
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
	var ingredient models.Ingredient
	if err := models.DB.First(&ingredient, ingredientID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingredient not found"})
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
	if err := models.DB.Model(&recipeIngredient).Updates(input).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	responseIngredient := RecipeIngredientsResponse{ID: int(ingredient.ID), Name: ingredient.Name, Portion: ingredient.Portion, Amount: recipeIngredient.Amount}
	c.JSON(http.StatusOK, gin.H{"data": responseIngredient})
}

// DELETE
func DeleteRecipeIngredient(c *gin.Context) {
	// Get records
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe not found"})
		return
	}
	var ingredient models.Ingredient
	if err := models.DB.First(&ingredient, c.Param("ingredientId")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ingredient not found"})
		return
	}
	// Delete recipe ingredient
	models.DB.Model(&recipe).Association("Ingredients").Delete(ingredient)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
