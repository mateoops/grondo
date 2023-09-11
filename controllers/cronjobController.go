package controllers

import (
	"grondo/db/sql"
	"grondo/models"

	"github.com/gin-gonic/gin"
)

// CronjobCreate godoc
// @Summary      Create Cronjob object
// @Tags         cronjob
// @Accept 		 json
// @Produce      json
// @Success      200  {array}  models.CronJob
// @Failure      500  {object}  string
// @Param        cronjob body models.CronJob true "CronJob object"
// @Router       /cronjob/ [post]
func CronjobCreate(c *gin.Context) {
	var body struct {
		Name      string
		Cron      string
		Enabled   bool
		Command   string
		Args      string
		MachineID uint
	}

	c.Bind(&body)

	job := models.CronJob{Name: body.Name, Cron: body.Cron, Enabled: body.Enabled, Command: body.Command, Args: body.Args, MachineID: body.MachineID}

	result := sql.DB.Create(&job)

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

// CronjobIndex godoc
// @Summary      List all CronJob objects
// @Tags         cronjob
// @Produce      json
// @Success      200  {object}  models.CronJob
// @Failure      500  {object}  string
// @Router       /cronjob [get]
func CronjobIndex(c *gin.Context) {
	var jobs []models.CronJob

	result := sql.DB.Find(&jobs)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": jobs,
	})
}

// CronjobShow godoc
// @Summary      Get Cronjob object by providing ID
// @Tags         cronjob
// @Produce      json
// @Success      200  {array}  models.CronJob
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /cronjob/{id} [get]
func CronjobShow(c *gin.Context) {
	id := c.Param("id")

	var job models.CronJob
	result := sql.DB.First(&job, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": job,
	})
}

// CronjobUpdate godoc
// @Summary      Update CronJob object by providing ID
// @Tags         cronjob
// @Accept 		 json
// @Produce      json
// @Success      200  {array}  models.CronJob
// @Failure      500  {object}  string
// @Param        cronjob body models.CronJob true "CronJob object"
// @Param        id  path  int  true  "ID"
// @Router       /cronjob/{id} [put]
func CronjobUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Name      string
		Cron      string
		Enabled   bool
		Command   string
		Args      string
		MachineId uint
	}

	c.Bind(&body)

	var job models.CronJob
	result := sql.DB.First(&job, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	sql.DB.Model(&job).Updates(models.CronJob{
		Name:      body.Name,
		Cron:      body.Cron,
		Enabled:   body.Enabled,
		Command:   body.Command,
		Args:      body.Args,
		MachineID: body.MachineId,
	})

	// Remove all scheduled jobs
	var schedule models.Schedule
	sql.DB.Delete(&schedule, "cron_job_id = ?", id)

	c.JSON(200, gin.H{
		"result": job,
	})
}

// CronjobDelete godoc
// @Summary      Delete CronJob object by providing ID
// @Tags         cronjob
// @Produce      json
// @Success      200  {object}  string
// @Failure      500  {object}  string
// @Param        id  path  int  true  "ID"
// @Router       /cronjob/{id} [delete]
func CronjobDelete(c *gin.Context) {
	id := c.Param("id")

	var job models.CronJob
	result := sql.DB.First(&job, id)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"message": result.Error,
		})
		return
	}

	sql.DB.Delete(&job)

	// Delete all scheduled jobs
	var schedule models.Schedule
	sql.DB.Delete(&schedule, "cron_job_id = ?", id)

	c.JSON(200, gin.H{
		"result": "success",
	})
}
