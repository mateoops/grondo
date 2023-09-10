package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// Create a new UserWithPassword credential object in the database
func UserPassCreate(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	c.Bind(&body)

	credential := models.UserWithPassword{Username: body.Username, Password: body.Password}

	result := sql.DB.Create(&credential)

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

// List all UserWithPassword credential objects in the database
func UserPassIndex(c *gin.Context) {
	var credentials []models.UserWithPassword

	result := sql.DB.Find(&credentials)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": credentials,
	})
}

// Show a specific UserWithPassword credential object in the database
func UserPassShow(c *gin.Context) {
	id := c.Param("id")

	var credential models.UserWithPassword

	result := sql.DB.First(&credential, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": credential,
	})
}

// Update a specific UserWithPassword credential object in the database
func UserPassUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Username string
		Password string
	}

	c.Bind(&body)

	credential := models.UserWithPassword{Username: body.Username, Password: body.Password}

	result := sql.DB.Model(&credential).Where("id = ?", id).Updates(&credential)

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

// Delete a specific UserWithPassword credential object in the database
func UserPassDelete(c *gin.Context) {
	id := c.Param("id")

	var credential models.UserWithPassword

	result := sql.DB.Delete(&credential, id)

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

// Create a new UserWithSSHKey credential object in the database
func UserSSHCreate(c *gin.Context) {
	var body struct {
		Username   string
		SSHKey     string
		Passphrase string
	}

	c.Bind(&body)

	credential := models.UserWithSSHKey{Username: body.Username, SSHKey: body.SSHKey, Passphrase: body.Passphrase}

	result := sql.DB.Create(&credential)

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

// List all UserWithSSHKey credential objects in the database
func UserSSHIndex(c *gin.Context) {
	var credentials []models.UserWithSSHKey

	result := sql.DB.Find(&credentials)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": credentials,
	})
}

// Show a specific UserWithSSHKey credential object in the database
func UserSSHShow(c *gin.Context) {
	id := c.Param("id")

	var credential models.UserWithSSHKey

	result := sql.DB.First(&credential, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": credential,
	})
}

// Update a specific UserWithSSHKey credential object in the database
func UserSSHUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Username   string
		SSHKey     string
		Passphrase string
	}

	c.Bind(&body)

	credential := models.UserWithSSHKey{Username: body.Username, SSHKey: body.SSHKey, Passphrase: body.Passphrase}

	result := sql.DB.Model(&credential).Where("id = ?", id).Updates(&credential)

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

// Delete a specific UserWithSSHKey credential object in the database
func UserSSHDelete(c *gin.Context) {
	id := c.Param("id")

	var credential models.UserWithSSHKey

	result := sql.DB.Delete(&credential, id)

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
