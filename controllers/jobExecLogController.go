package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// JobExecLogIndex godoc
// @Summary      List all JobExecLog objects
// @Tags         cronjob execution_log
// @Produce      json
// @Success      200  {object}  models.JobExecLog
// @Failure      500  {object}  string
// @Router       /log [get]
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

// JobExecLogShow godoc
// @Summary      Get JobExecLog object by providing ID
// @Tags         cronjob execution_log
// @Produce      json
// @Success      200  {array}  models.JobExecLog
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /log/{id} [get]
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
