package migrate

import (
	"grondo/db/sql"
	"grondo/models"
	"log"
)

func init() {
	sql.InitDB()
}

func MigrateDB() {
	log.Println("Migrating database...")

	sql.DB.AutoMigrate(&models.CronJob{})
	sql.DB.AutoMigrate(&models.Schedule{})
	sql.DB.AutoMigrate(&models.Queue{})
	sql.DB.AutoMigrate(&models.JobExecLog{})
	sql.DB.AutoMigrate(&models.Machine{})
	sql.DB.AutoMigrate(&models.UserWithPassword{})
	sql.DB.AutoMigrate(&models.UserWithSSHKey{})

	log.Println("Migration completed!")
}
