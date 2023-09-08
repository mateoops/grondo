package models

import (
	"gorm.io/gorm"
)

type Queue struct {
	gorm.Model

	JobExecLogID uint

	ScheduleID uint
	Schedule   Schedule
}
