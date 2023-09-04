package scheduler

import (
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

func ParseCron(cronjobString string) string {
	cronParser := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
	schedule, err := cronParser.Parse(cronjobString)
	if err != nil {
		log.Fatalln("Error parsing crontab expression:", err)
	}

	now := time.Now()
	numTimes := 5
	nextTimes := make([]time.Time, numTimes)

	for i := 0; i < numTimes; i++ {
		nextTimes[i] = schedule.Next(now)
		now = nextTimes[i]
	}

	for i, t := range nextTimes {
		fmt.Printf("Next time #%d: %s\n", i+1, t.Format(time.RFC3339))
		fmt.Println(nextTimes)
	}
	return "x"
}
