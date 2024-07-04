package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"recipiary/controllers"
	"recipiary/models"
)

func main() {
	r := gin.Default()
	models.Connect()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/categories", controllers.IndexCategories)
	r.GET("/categories/:id", controllers.GetCategory)
	r.POST("/categories", controllers.CreateCategory)
	r.PATCH("/categories/:id", controllers.UpdateCategory)
	r.DELETE("/categories/:id", controllers.DeleteCategory)

	r.GET("/ingredients", controllers.IndexIngredients)
	r.GET("/ingredients/:id", controllers.GetIngredient)
	r.POST("/ingredients", controllers.CreateIngredient)
	r.PATCH("/ingredients/:id", controllers.UpdateIngredient)
	r.DELETE("/ingredients/:id", controllers.DeleteIngredient)

	r.Run() // listen and serve on 0.0.0.0:8080
}
