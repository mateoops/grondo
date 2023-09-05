package controllers

import (
	sql "grondo/db"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

func CronjobCreate(c *gin.Context) {
	var body struct {
		Name    string
		Cron    string
		Enabled bool
		Command string
	}

	c.Bind(&body)

	post := models.CronJob{Name: body.Name, Cron: body.Cron, Enabled: body.Enabled, Command: body.Command}

	result := sql.DB.Create(&post)

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

func CronjobIndex(c *gin.Context) {
	var posts []models.CronJob

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

func CronjobShow(c *gin.Context) {
	id := c.Param("id")

	var post models.CronJob
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

func CronjobUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name    string
		Cron    string
		Enabled bool
		Command string
	}

	c.Bind(&body)

	var post models.CronJob
	result := sql.DB.First(&post, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	sql.DB.Model(&post).Updates(models.CronJob{
		Name:    body.Name,
		Cron:    body.Cron,
		Enabled: body.Enabled,
		Command: body.Command,
	})

	c.JSON(200, gin.H{
		"result": post,
	})
}

func CronjobDelete(c *gin.Context) {
	id := c.Param("id")

	var post models.CronJob
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
