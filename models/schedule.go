package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model

	NextOccur time.Time
	Status    string
	CronJobID uint
	CronJob   CronJob
}
