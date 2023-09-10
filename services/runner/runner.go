package runner

import (
	"grondo/db/sql"
	"grondo/models"
	"grondo/utils/sshutil"
	"log"
	"os/exec"
	"time"
)

func finishJob(status string, output string, q models.Schedule, startTime time.Time, endTime time.Time) {

	sql.DB.Create(&models.JobExecLog{ScheduleID: q.ID, Status: status, Output: output, StartTime: startTime, EndTime: endTime})
	sql.DB.Where("id = ?", q.ID).First(&q).Update("status", "FINISHED")
	log.Printf("Finished job: %v", q.CronJob.Name)

}

func runCommandOnLocalMachine(q models.Schedule) {

	sql.DB.Where("id = ?", q.ID).First(&q).Update("status", "STARTED")
	startTime := time.Now()

	// Create command object
	var cmd *exec.Cmd
	if q.CronJob.Args == "" {
		cmd = exec.Command(q.CronJob.Command)
	} else {
		cmd = exec.Command(q.CronJob.Command, q.CronJob.Args)
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

func runCommandOnRemoteMachine(q models.Schedule) {

	sql.DB.Where("id = ?", q.ID).First(&q).Update("status", "STARTED")
	startTime := time.Now()

	var username string
	var password string
	if q.CronJob.Machine.UserWithSSHKeyID > 0 {
		username = q.CronJob.Machine.UserWithSSHKey.Username
		password = q.CronJob.Machine.UserWithSSHKey.SSHKey
	} else if q.CronJob.Machine.UserWithPasswordID > 0 {
		username = q.CronJob.Machine.UserWithPassword.Username
		password = q.CronJob.Machine.UserWithPassword.Password
	} else {
		finishJob("ERROR", "Machine hasn't credentials!", q, startTime, startTime)
		return
	}
	status, logText := sshutil.RunCommandOverSSH(q.CronJob.Machine.Address, q.CronJob.Machine.Port, username, password, q.CronJob.Command+" "+q.CronJob.Args)

	endTime := time.Now()
	finishJob(status, logText, q, startTime, endTime)
	return
}

func StartRunner() {
	log.Println("Runner started...")
	for {

		var queue []models.Queue
		sql.DB.Preload("Schedule.CronJob.Machine.UserWithSSHKey").Preload("Schedule.CronJob.Machine.UserWithPassword").Find(&queue)
		for _, q := range queue {
			queueCopy := q
			log.Printf("Started job: %v", queueCopy.Schedule.CronJob.Name)
			if queueCopy.Schedule.CronJob.Machine.Address == "localhost" || queueCopy.Schedule.CronJob.Machine.Address == "127.0.0.1" {
				go runCommandOnLocalMachine(queueCopy.Schedule)
			} else {
				go runCommandOnRemoteMachine(queueCopy.Schedule)
			}
			sql.DB.Delete(&queueCopy)
		}
		time.Sleep(2 * time.Second)
	}
}
