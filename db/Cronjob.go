package db

type CronJob struct {
	Id      int64
	Name    string
	Cron    string
	Enabled bool
	Command string
}
