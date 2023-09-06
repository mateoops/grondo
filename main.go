package main

import (
	"grondo/db/sql"
	"grondo/routers"
	"grondo/services/queueManager"
	"grondo/services/scheduler"
	"grondo/utils"
)

func init() {
	utils.LoadYamlConfig()
	sql.InitDB()

	go scheduler.StartScheduler()
	go queueManager.StartQueueManager()
}

func main() {
	r := routers.InitRouter()
	r.Run()
}
