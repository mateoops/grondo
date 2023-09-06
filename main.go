package main

import (
	"fmt"
	"grondo/controllers"
	"grondo/db/sql"
	"grondo/services/queueManager"
	"grondo/services/scheduler"
	"grondo/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadYamlConfig()
	sql.InitDB()

	go scheduler.StartScheduler()
	go queueManager.StartQueueManager()
}

func main() {
	log.Println("Config:")
	log.Println(utils.AppConfig)

	contextPath := fmt.Sprintf("%s/%s", utils.AppConfig.Server.ContextPath, utils.AppConfig.Server.ApiVersion)

	r := gin.Default()

	r.POST(contextPath+"/cronjob", controllers.CronjobCreate)
	r.GET(contextPath+"/cronjob", controllers.CronjobIndex)
	r.GET(contextPath+"/cronjob/:id", controllers.CronjobShow)
	r.PUT(contextPath+"/cronjob/:id", controllers.CronjobUpdate)
	r.DELETE(contextPath+"/cronjob/:id", controllers.CronjobDelete)

	r.GET(contextPath+"/schedule", controllers.ScheduleIndex)
	r.GET(contextPath+"/schedule/:id", controllers.ScheduleShow)

	r.GET(contextPath+"/queue", controllers.CronjobQueue)
	r.GET(contextPath+"/queue/:id", controllers.CronjobQueueShow)

	r.Run()
}
