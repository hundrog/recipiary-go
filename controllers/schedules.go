package controllers

import (
	"net/http"
	"recipiary/models"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateScheduleInput struct {
	StartDate time.Time `binding:"required,ltefield=FinalDate" time_format:"2006-01-02"`
	FinalDate time.Time `binding:"required" time_format:"2006-01-02"`
}

type UpdateScheduleInput struct {
	FromDate string
	ToDate   string
}

// INDEX
func IndexSchedules(c *gin.Context) {
	var schedules []models.Schedule
	models.DB.Preload("Recipes").Find(&schedules)

	c.JSON(http.StatusOK, gin.H{"data": schedules})
}

// POST
func CreateSchedule(c *gin.Context) {
	// Validate input
	var input CreateScheduleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create schedule
	schedule := models.Schedule{StartDate: input.StartDate, FinalDate: input.FinalDate}
	models.DB.Create(&schedule)

	c.JSON(http.StatusOK, gin.H{"data": schedule})
}
