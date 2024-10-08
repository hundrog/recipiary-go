package main

import (
	"net/http"
	"os"
	"recipiary/auth"
	"recipiary/controllers"
	"recipiary/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func main() {
	models.Connect()

	auth.Init()
	r := gin.Default()
	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{os.Getenv("SUPERTOKENS_WEB_DOMAIN")},
		AllowMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowHeaders: append([]string{"content-type"},
			supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Use(auth.SuperTokens())
	r.Use(auth.VerifySession())

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
	r.PATCH("/recipes/:id/ingredients/:ingredientId", controllers.UpdateRecipeIngredient)
	r.DELETE("/recipes/:id/ingredients/:ingredientId", controllers.DeleteRecipeIngredient)

	r.GET("/recipes/:id/instructions", controllers.IndexInstructions)
	r.POST("/recipes/:id/instructions", controllers.CreateInstruction)
	r.POST("/recipes/:id/instructions_bulk", controllers.CreateInstructionBulk)
	r.PATCH("/recipes/:id/instructions/:instructionId", controllers.UpdateInstruction)
	r.PATCH("/recipes/:id/instructions/position", controllers.UpdateInstructionsOrder)
	r.DELETE("/recipes/:id/instructions/:instructionId", controllers.DeleteInstruction)

	r.GET("/schedules", controllers.IndexSchedules)
	r.POST("/schedules", controllers.CreateSchedule)
	r.GET("/schedules/:id", controllers.GetSchedule)
	r.PATCH("/schedules/:id", controllers.UpdateSchedule)
	r.DELETE("/schedules/:id", controllers.DeleteSchedule)

	r.POST("/schedules/:id/recipes", controllers.CreateScheduleRecipes)
	r.PATCH("/schedules/:id/recipes", controllers.UpdateScheduleRecipes)
	r.DELETE("/schedules/:id/recipes", controllers.DeleteScheduleRecipes)
	r.DELETE("/schedules/:id/clear", controllers.ClearScheduleRecipes)

	r.Run()
}
