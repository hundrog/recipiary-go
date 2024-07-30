package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recipiary/models"
)

type CreateInstructionInput struct {
	Content string `binding:"required"`
}

type CreateInstructionBulkInput struct {
	Content []string `binding:"required"`
}
type UpdateInstructionInput struct {
	Content string
}

type IngredientsResponse struct {
	ID      int
	Name    string
	Portion string
	Amount  int
}

// INDEX
func IndexInstructions(c *gin.Context) {
	var instructions []models.Instruction
	models.DB.Where("recipe_id = ?", c.Param("id")).Find(&instructions)

	c.JSON(http.StatusOK, gin.H{"data": instructions})
}

// POST
func CreateInstruction(c *gin.Context) {
	// Validate imput
	var input CreateInstructionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Retrieve recipe
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe doesn't exists"})
		return
	}
	// Create ingredient
	instruction := models.Instruction{RecipeID: int(recipe.ID), Content: input.Content}
	models.DB.Create(&instruction)

	c.JSON(http.StatusOK, gin.H{"data": instruction})
}

// PATCH
func UpdateInstruction(c *gin.Context) {
	// Validate imput
	var input UpdateInstructionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Retrieve recipe
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe doesn't exists"})
		return
	}
	// Retrieve Instruction
	var instruction models.Instruction
	if err := models.DB.First(&instruction, c.Param("instruction_id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Instruction doesn't exists"})
		return
	}
	// Create ingredient
	models.DB.Model(&instruction).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": instruction})
}

// POST
func CreateInstructionBulk(c *gin.Context) {
	// Validate imput
	var input CreateInstructionBulkInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Retrieve recipe
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe doesn't exists"})
		return
	}

	// Create ingredients
	var instructions []models.Instruction
	for _, element := range input.Content {
		instruction := models.Instruction{RecipeID: int(recipe.ID), Content: element}
		instructions = append(instructions, instruction)
	}

	models.DB.Create(&instructions)

	c.JSON(http.StatusOK, gin.H{"data": instructions})
}

// DELETE
func DeleteInstruction(c *gin.Context) {
	// Retrieve recipe
	var recipe models.Recipe
	if err := models.DB.First(&recipe, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Recipe doesn't exists"})
		return
	}

	var instruction models.Instruction
	if err := models.DB.First(&instruction, c.Param("instructionId")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Instruction doesn't exists"})
		return
	}

	models.DB.Delete(&instruction)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
