package models

import (
	"time"

	"gorm.io/gorm"
)

type CronJobNextOccur struct {
	gorm.Model
	CronJobID uint
	CronJob   CronJob
	NextOccur time.Time
}
