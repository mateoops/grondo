package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// Show all exec log objects in the database
func JobExecLogIndex(c *gin.Context) {
	var e []models.JobExecLog

	result := sql.DB.Preload("Schedule").Find(&e)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": e,
	})
}

// Show a specific queue object in the database
func JobExecLogShow(c *gin.Context) {
	id := c.Param("id")

	var e models.JobExecLog
	result := sql.DB.Preload("Schedule").First(&e, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": e,
	})
}
