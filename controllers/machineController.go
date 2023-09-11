package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// MachineCreate godoc
// @Summary      Create Machine object
// @Tags         cronjob machine
// @Accept 		 json
// @Produce      json
// @Success      200  {array}  models.Machine
// @Failure      500  {object}  string
// @Param        machine body models.Machine true "Machine object"
// @Router       /machine [post]
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

// MachineIndex godoc
// @Summary      List all Machine objects
// @Tags         cronjob machine
// @Produce      json
// @Success      200  {object}  models.Machine
// @Failure      500  {object}  string
// @Router       /machine [get]
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

// MachineShow godoc
// @Summary      Get Machine object by providing ID
// @Tags         cronjob machine
// @Produce      json
// @Success      200  {array}  models.Machine
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /machine/{id} [get]
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

// MachineUpdate godoc
// @Summary      Update Machine object by providing ID
// @Tags         cronjob machine
// @Accept 		 json
// @Produce      json
// @Success      200  {array}  models.Machine
// @Failure      500  {object}  string
// @Param        machine body models.Machine true "Machine object"
// @Param        id  path  int  true  "ID"
// @Router       /machine/{id} [put]
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

// MachineDelete godoc
// @Summary      Delete Machine object by providing ID
// @Tags         cronjob machine
// @Produce      json
// @Success      200  {object}  string
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /machine/{id} [delete]
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
