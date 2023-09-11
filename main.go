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

// @title           GRONDO API
// @version         0.0.1
// @description     Centralized cron jobs with web GUI.
// @termsOfService  https://github.com/mateoops/grondo

// @contact.name   Author
// @contact.url    https://github.com/mateoops/grondo
// @contact.email  mateoops.dev@gmail.com

// @license.name  MIT License
// @license.url   https://github.com/mateoops/grondo/blob/main/LICENSE

// @host      localhost:8080
// @BasePath  /api/v1

func main() {
	r := routers.InitRouter()
	r.Run()
}
