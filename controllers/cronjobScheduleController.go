package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// Show all schedule objects in the database
func ScheduleIndex(c *gin.Context) {
	var schedules []models.Schedule

	result := sql.DB.Preload("CronJob").Find(&schedules)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": schedules,
	})
}

// List all schedule objects in the database
func ScheduleShow(c *gin.Context) {
	id := c.Param("id")

	var schedule models.Schedule
	result := sql.DB.Preload("CronJob").First(&schedule, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": schedule,
	})
}
