package main

import (
	"grondo/db/sql"
	"grondo/models"
	"log"
)

func init() {
	sql.InitDB()
}

func main() {
	log.Println("Migrating database...")

	sql.DB.AutoMigrate(&models.CronJob{})
	sql.DB.AutoMigrate(&models.Schedule{})
	sql.DB.AutoMigrate(&models.Queue{})

	log.Println("Migration completed!")
}
