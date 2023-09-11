package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// CronjobQueueIndex godoc
// @Summary      List all Cronjob Queue objects
// @Tags         cronjob queue
// @Produce      json
// @Success      200  {object}  models.Queue
// @Failure      500  {object}  string
// @Router       /queue [get]
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

// CronjobQueueShow godoc
// @Summary      Get Cronjob Queue object by providing ID
// @Tags         cronjob queue
// @Produce      json
// @Success      200  {array}  models.Queue
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /queue/{id} [get]
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
