package controllers

import (
	"net/http"
	"recipiary/models"

	"github.com/gin-gonic/gin"
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

type UpdatePosition struct {
	ID       int  `binding:"required"`
	Position *int `binding:"required"`
}

type UpdatePositionPayload struct {
	Updates []UpdatePosition `binding:"dive"`
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
	models.DB.Where("recipe_id = ?", c.Param("id")).Order("position").Find(&instructions)

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
	if err := models.DB.First(&instruction, c.Param("instructionId")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Instruction doesn't exists"})
		return
	}
	// Update ingredient
	models.DB.Model(&instruction).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": instruction})
}

func UpdateInstructionsOrder(c *gin.Context) {
	var input UpdatePositionPayload
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, element := range input.Updates {
		_ = i
		var instruction models.Instruction
		if err := models.DB.First(&instruction, element.ID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Instruction doesn't exists"})
			return
		}

		models.DB.Model(&instruction).Update("position", element.Position)
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
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
