package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"recipiary/controllers"
	"recipiary/models"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	models.Connect()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/categories", controllers.IndexCategories)
	r.POST("/categories", controllers.CreateCategory)
	r.GET("/categories/:id", controllers.GetCategory)
	r.PATCH("/categories/:id", controllers.UpdateCategory)
	r.DELETE("/categories/:id", controllers.DeleteCategory)

	r.GET("/ingredients", controllers.IndexIngredients)
	r.POST("/ingredients", controllers.CreateIngredient)
	r.GET("/ingredients/:id", controllers.GetIngredient)
	r.PATCH("/ingredients/:id", controllers.UpdateIngredient)
	r.DELETE("/ingredients/:id", controllers.DeleteIngredient)

	r.GET("/recipes", controllers.IndexRecipes)
	r.POST("/recipes", controllers.CreateRecipe)
	r.GET("/recipes/:id", controllers.GetRecipe)
	r.PATCH("/recipes/:id", controllers.UpdateRecipe)
	r.DELETE("/recipes/:id", controllers.DeleteRecipe)

	r.GET("/recipes/:id/ingredients", controllers.GetRecipeIngredients)
	r.POST("/recipes/:id/ingredients", controllers.InsertRecipeIngredient)
	r.PATCH("/recipes/:id/ingredients/:ingredientID", controllers.UpdateRecipeIngredient)
	r.DELETE("/recipes/:id/ingredients/:ingredientID", controllers.DeleteRecipeIngredient)

	r.GET("/recipes/:id/instructions", controllers.IndexInstructions)
	r.POST("/recipes/:id/instructions", controllers.CreateInstruction)
	r.POST("/recipes/:id/instructions_bulk", controllers.CreateInstructionBulk)

	r.Run() // listen and serve on 0.0.0.0:8080
}
