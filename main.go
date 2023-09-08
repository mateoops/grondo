package main

import (
	"grondo/db/sql"
	"grondo/routers"
	"grondo/services/queueManager"
	"grondo/services/runner"
	"grondo/services/scheduler"
	"grondo/utils"
)

func init() {
	utils.LoadYamlConfig()
	sql.InitDB()

	go scheduler.StartScheduler()
	go queueManager.StartQueueManager()
	go runner.StartRunner()
}

func main() {
	r := routers.InitRouter()
	r.Run()
}
