package scheduler

import (
	"grondo/db/sql"
	"grondo/models"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

// Translate a crontab expression into a time object
// and return one occurrence or more if the occurrences will be in the next X minutes
func getNearestOccurs(cronjobString string, minutes uint) []time.Time {
	cronParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	schedule, err := cronParser.Parse(cronjobString)
	if err != nil {
		log.Fatalln("Error parsing crontab expression:", err)
	}

	// Calculate next occurrence
	now := time.Now()
	to := now.Add(time.Duration(minutes) * time.Minute)
	nextTimes := []time.Time{}

	// Always add the first occurrence
	nextTime := schedule.Next(now)
	nextTimes = append(nextTimes, nextTime)

	for {
		if nextTime.After(to) {
			break
		} else {
			nextTimes = append(nextTimes, nextTime)
			nextTime = schedule.Next(nextTime)
		}
	}

	return nextTimes
}

// Clean up queued jobs that are older than 1 minutes (missed) by setting their status to MISSED QUEUE
// in the scheduler database and remove them from the queue
func schedulerCleaner() {
	var queuedJobs []models.Schedule
	sql.DB.Find(&queuedJobs)

	for _, job := range queuedJobs {
		// Missing is when the job is scheduled but not queued since 1 minute
		if job.Status == "SCHEDULED" && job.NextOccur.Before(time.Now().Add(-1*time.Minute)) {
			jobCopy := job
			jobCopy.Status = "MISSED SCHEDULE"
			sql.DB.Save(&jobCopy)
			log.Println("Missed scheduled job (job not queued on time):", jobCopy.ID)
		}
	}
}

// StartScheduler starts the scheduler. Process which checks for jobs that need to be queued and queues them.
// Checks every 15 seconds.
func StartScheduler() {
	for {
		log.Println("Scheduler started...")
		schedulerCleaner()

		var jobs []models.CronJob
		sql.DB.Find(&jobs)

		for _, job := range jobs {
			if job.Enabled {
				// Get occurrences for the next 5 minutes (or only one if it isn't in the next 5 minutes)
				occurs := getNearestOccurs(job.Cron, 5)
				for _, occur := range occurs {
					// Check if the job is already scheduled
					if sql.DB.Where("cron_job_id = ? AND next_occur = ?", job.ID, occur).Find(&models.Schedule{}).RowsAffected == 0 {
						oc := models.Schedule{CronJobID: job.ID, NextOccur: occur, Status: "SCHEDULED"}
						sql.DB.Create(&oc)
						log.Println("Planed next occur of job:", job.Name)
					}
				}
			}
		}
		time.Sleep(15 * time.Second)
	}
}
