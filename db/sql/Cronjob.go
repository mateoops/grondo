package sql

import (
	"grondo/db"
	"grondo/db/sql"
	"log"
)

func CreateCronJob(cronjob db.CronJob) (newCronjob db.CronJob, err error) {
	cronjob.Id = 1
	cronjob.Name = "test"
	cronjob.Cron = "0 0 0 0 0"
	cronjob.Enabled = true
	cronjob.Command = "test"

	newCronjob = cronjob

	sql.CreateCronJobTable()

	insertStmt, err := SqlDb.Prepare("INSERT INTO users(username, email) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer insertStmt.Close()

	return
}
