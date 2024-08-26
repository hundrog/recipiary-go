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
	StartDate time.Time `binding:"ltefield=FinalDate" time_format:"2006-01-02"`
	FinalDate time.Time `time_format:"2006-01-02"`
}

// INDEX
func IndexSchedules(c *gin.Context) {
	var schedules []models.Schedule
	today := time.Now().Format("2006-01-02")
	models.DB.Preload("Recipes").Order("start_date").Where("final_date > ?", today).Find(&schedules)

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

func GetSchedule(c *gin.Context) {
	// Get record
	var schedule models.Schedule
	if err := models.DB.Preload("Recipes").First(&schedule, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": schedule})
}

func UpdateSchedule(c *gin.Context) {
	// Get record
	var schedule models.Schedule
	if err := models.DB.First(&schedule, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	// Validate input
	var input UpdateScheduleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update schedule
	models.DB.Model(&schedule).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": schedule})
}

func DeleteSchedule(c *gin.Context) {
	// Get record
	var schedule models.Schedule
	if err := models.DB.First(&schedule, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}
	// Delete category
	models.DB.Delete(&schedule)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
