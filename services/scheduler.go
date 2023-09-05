package scheduler

import (
	"fmt"
	sql "grondo/db"
	"grondo/models"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func getNearestOccurs(cronjobString string, minutes uint) []time.Time {
	cronParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	schedule, err := cronParser.Parse(cronjobString)
	if err != nil {
		log.Fatalln("Error parsing crontab expression:", err)
	}

	now := time.Now()
	to := now.Add(15 * time.Minute)
	nextTimes := []time.Time{}

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

func StartScheduler() {
	for {
		fmt.Println("Scheduler running...")

		var posts []models.CronJob
		sql.DB.Find(&posts)

		for _, post := range posts {
			if post.Enabled {
				occurs := getNearestOccurs(post.Cron, 15)
				for _, occur := range occurs {
					if sql.DB.Where("cron_job_id = ? AND next_occur = ?", post.ID, occur).Find(&models.CronJobNextOccur{}).RowsAffected == 0 {
						oc := models.CronJobNextOccur{CronJobID: post.ID, NextOccur: occur}
						sql.DB.Create(&oc)
					}
				}
			}
		}
		time.Sleep(15 * time.Second)
	}
}
