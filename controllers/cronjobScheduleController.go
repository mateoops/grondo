package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// ScheduleIndex godoc
// @Summary      List all Schedule objects
// @Tags         cronjob schedule
// @Produce      json
// @Success      200  {object}  models.Schedule
// @Failure      500  {object}  string
// @Router       /schedule [get]
func ScheduleIndex(c *gin.Context) {
	var schedules []models.Schedule

	result := sql.DB.Preload("CronJob").Find(&schedules)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": schedules,
	})
}

// ScheduleShow godoc
// @Summary      Get Schedule object by providing ID
// @Tags         cronjob schedule
// @Produce      json
// @Success      200  {array}  models.Schedule
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /schedule/{id} [get]
func ScheduleShow(c *gin.Context) {
	id := c.Param("id")

	var schedule models.Schedule
	result := sql.DB.Preload("CronJob").First(&schedule, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": schedule,
	})
}
