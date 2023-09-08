package runner

import (
	"grondo/db/sql"
	"grondo/models"
	"log"
	"os/exec"
	"time"
)

func finishJob(status string, output string, q models.Queue, startTime time.Time, endTime time.Time) {

	sql.DB.Create(&models.JobExecLog{ScheduleID: q.Schedule.ID, Status: status, Output: output, StartTime: startTime, EndTime: endTime})
	sql.DB.Where("id = ?", q.ScheduleID).First(&q.Schedule).Update("status", "FINISHED")
	sql.DB.Delete(&q)
	log.Printf("Finished job: %v", q.Schedule.CronJob.Name)

}

func runCommandOnLocalMachine(q models.Queue) {

	sql.DB.Where("id = ?", q.ScheduleID).First(&q.Schedule).Update("status", "STARTED")
	startTime := time.Now()

	// Create command object
	var cmd *exec.Cmd
	if q.Schedule.CronJob.Args == "" {
		cmd = exec.Command(q.Schedule.CronJob.Command)
	} else {
		cmd = exec.Command(q.Schedule.CronJob.Command, q.Schedule.CronJob.Args)
	}

	output, err := cmd.CombinedOutput()

	var logText string
	var status string
	if err != nil {
		logText = "Error running command on local machine. Details:" + err.Error() + "\n" + string(output)
		status = "ERROR"
	} else {
		logText = string(output)
		status = "SUCCESS"
	}
	endTime := time.Now()
	finishJob(status, logText, q, startTime, endTime)
}

func StartRunner() {
	log.Println("Runner started...")
	for {

		var queue []models.Queue
		sql.DB.Preload("Schedule.CronJob").Find(&queue)

		for _, q := range queue {
			queueCopy := q
			log.Printf("Started job: %v", queueCopy.Schedule.CronJob.Name)
			go runCommandOnLocalMachine(queueCopy)
		}
		time.Sleep(2 * time.Second)
	}
}
