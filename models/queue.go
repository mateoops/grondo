package models

import (
	"gorm.io/gorm"
)

type Queue struct {
	gorm.Model

	ScheduleID uint
	Schedule   Schedule
}
