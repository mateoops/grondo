package models

import (
	"gorm.io/gorm"
)

type CronJob struct {
	gorm.Model
	Name    string
	Cron    string
	Enabled bool
	Command string
	Args    string
}
