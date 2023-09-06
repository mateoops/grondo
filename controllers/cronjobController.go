package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// Create a new cronjob object in the database
func CronjobCreate(c *gin.Context) {
	var body struct {
		Name    string
		Cron    string
		Enabled bool
		Command string
	}

	c.Bind(&body)

	job := models.CronJob{Name: body.Name, Cron: body.Cron, Enabled: body.Enabled, Command: body.Command}

	result := sql.DB.Create(&job)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": "success",
	})
}

// List all cronjob objects in the database
func CronjobIndex(c *gin.Context) {
	var jobs []models.CronJob

	result := sql.DB.Find(&jobs)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": jobs,
	})
}

// Show a specific cronjob object in the database
func CronjobShow(c *gin.Context) {
	id := c.Param("id")

	var job models.CronJob
	result := sql.DB.First(&job, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": job,
	})
}

// Update a specific cronjob object in the database
func CronjobUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name    string
		Cron    string
		Enabled bool
		Command string
	}

	c.Bind(&body)

	var job models.CronJob
	result := sql.DB.First(&job, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	sql.DB.Model(&job).Updates(models.CronJob{
		Name:    body.Name,
		Cron:    body.Cron,
		Enabled: body.Enabled,
		Command: body.Command,
	})

	// Remove all scheduled jobs
	var schedule models.Schedule
	sql.DB.Delete(&schedule, "cron_job_id = ?", id)

	c.JSON(200, gin.H{
		"result": job,
	})
}

// Delete a specific cronjob object in the database
func CronjobDelete(c *gin.Context) {
	id := c.Param("id")

	var job models.CronJob
	result := sql.DB.First(&job, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	sql.DB.Delete(&job)

	// Delete all scheduled jobs
	var schedule models.Schedule
	sql.DB.Delete(&schedule, "cron_job_id = ?", id)

	c.JSON(200, gin.H{
		"result": "success",
	})
}
