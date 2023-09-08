package models

import (
	"time"

	"gorm.io/gorm"
)

type JobExecLog struct {
	gorm.Model

	Status     string
	Output     string
	StartTime  time.Time
	EndTime    time.Time
	ScheduleID uint
	Schedule   Schedule
}
