package controllers

import (
	sql "grondo/db"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

func CronjobNextOccurIndex(c *gin.Context) {
	var posts []models.CronJobNextOccur

	result := sql.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": posts,
	})
}

func CronjobNextOccurShow(c *gin.Context) {
	id := c.Param("id")

	var post models.CronJobNextOccur
	result := sql.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": post,
	})
}

func CronjobNextOccurDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.CronJobNextOccur
	result := sql.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	sql.DB.Delete(&post)

	c.JSON(200, gin.H{
		"result": "success",
	})
}
