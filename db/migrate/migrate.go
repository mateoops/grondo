package main

import (
	sql "grondo/db"
	"grondo/models"
)

func init() {
	sql.InitDB()
}

func main() {
	sql.DB.AutoMigrate(&models.CronJob{})
	sql.DB.AutoMigrate(&models.CronJobNextOccur{})
}
