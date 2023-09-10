package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// Create a new machine object in the database
func MachineCreate(c *gin.Context) {
	var body struct {
		Address            string
		Port               uint
		UserWithPasswordID uint
		UserWithSSHKeyID   uint
	}

	c.Bind(&body)

	machine := models.Machine{Address: body.Address, Port: body.Port, UserWithPasswordID: body.UserWithPasswordID, UserWithSSHKeyID: body.UserWithSSHKeyID}

	result := sql.DB.Create(&machine)

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

// List all machine objects in the database
func MachineIndex(c *gin.Context) {
	var machines []models.Machine

	result := sql.DB.Find(&machines)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": machines,
	})
}

// Show a specific machine object in the database
func MachineShow(c *gin.Context) {
	id := c.Param("id")

	var machine models.Machine

	result := sql.DB.First(&machine, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": machine,
	})
}

// Update a specific machine object in the database
func MachineUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Address            string
		Port               uint
		UserWithPasswordID uint
		UserWithSSHKeyID   uint
	}

	c.Bind(&body)

	var machine models.Machine
	result := sql.DB.First(&machine, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	sql.DB.Model(&machine).Updates(models.Machine{
		Address:            body.Address,
		Port:               body.Port,
		UserWithPasswordID: body.UserWithPasswordID,
		UserWithSSHKeyID:   body.UserWithSSHKeyID,
	})

	c.JSON(200, gin.H{
		"result": "success",
	})
}

// Delete a specific machine object in the database
func MachineDelete(c *gin.Context) {
	id := c.Param("id")

	var machine models.Machine

	result := sql.DB.Delete(&machine, id)

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
