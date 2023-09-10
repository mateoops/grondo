package main

import (
	"grondo/db/migrate"
	"grondo/db/sql"
	"grondo/routers"
	"grondo/services/queueManager"
	"grondo/services/runner"
	"grondo/services/scheduler"
	"grondo/utils/config"
)

func init() {
	config.LoadYamlConfig()
	migrate.MigrateDB()
	sql.InitDB()

	go scheduler.StartScheduler()
	go queueManager.StartQueueManager()
	go runner.StartRunner()
}

func main() {
	r := routers.InitRouter()
	r.Run()
}
