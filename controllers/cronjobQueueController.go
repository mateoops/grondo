package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// List all queue objects in the database
func CronjobQueueIndex(c *gin.Context) {
	var queueElements []models.Queue

	result := sql.DB.Preload("Schedule").Find(&queueElements)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": queueElements,
	})
}

// Show a specific queue object in the database
func CronjobQueueShow(c *gin.Context) {
	id := c.Param("id")

	var queue models.Queue
	result := sql.DB.Preload("Schedule").First(&queue, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": queue,
	})
}
