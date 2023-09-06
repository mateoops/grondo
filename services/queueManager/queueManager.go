package queueManager

import (
	"grondo/db/sql"
	"grondo/models"
	"log"
	"time"
)

// Check if a job should be run based on the next occur time (within 1 minute)
func checkJobShouldBeRun(job models.Schedule) bool {
	now := time.Now()
	if job.NextOccur.Before(now) && job.NextOccur.After(now.Add(-1*time.Minute)) {
		return true
	} else {
		return false
	}
}

// Clean up queued jobs that are older than 1 minutes (missed) by setting their status to MISSED QUEUE
func queueCleaner() {
	var queuedJobs []models.Queue
	sql.DB.Find(&queuedJobs)

	for _, job := range queuedJobs {
		jobCopy := job
		if jobCopy.CreatedAt.Before(time.Now().Add(-1 * time.Minute)) {
			var Schedule models.Schedule
			sql.DB.First(&Schedule, job.ScheduleID)

			Schedule.Status = "MISSED QUEUE"
			sql.DB.Save(Schedule)

			sql.DB.Delete(&jobCopy)
			log.Println("Missed queued job (job not started on time):", jobCopy.ID)
		}
	}
}

// StartQueueManager starts the queue manager. Process which checks for jobs that need to be queued and queues them.
// Checks every 20 seconds.
func StartQueueManager() {
	for {
		log.Println("QueueManager started...")

		queueCleaner()
		// Check if there are any cron jobs that need to be queued
		var jobs []models.Schedule
		sql.DB.Find(&jobs)

		for _, job := range jobs {
			// Check if the job is scheduled and if it should be run
			if job.Status == "SCHEDULED" && checkJobShouldBeRun(job) {
				jobCopy := job
				// Create a new object in queue database
				q := models.Queue{ScheduleID: jobCopy.ID}
				sql.DB.Create(&q)
				// Update the status of the job to QUEUED
				jobCopy.Status = "QUEUED"
				sql.DB.Save(&jobCopy)
				log.Println("Queued job:", jobCopy.ID)
			}
		}
		// sleep 5 seconds
		time.Sleep(5 * time.Second)
	}

}
