package main

import (
	"fmt"
	"grondo/controllers"
	sql "grondo/db"
	scheduler "grondo/services"
	"grondo/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadYamlConfig()
	sql.InitDB()

	go scheduler.StartScheduler()
}

func main() {
	fmt.Println("Config:")
	fmt.Println(utils.AppConfig)

	contextPath := fmt.Sprintf("%s/%s", utils.AppConfig.Server.ContextPath, utils.AppConfig.Server.ApiVersion)

	r := gin.Default()

	r.POST(contextPath+"/cronjob", controllers.CronjobCreate)
	r.GET(contextPath+"/cronjob", controllers.CronjobIndex)
	r.GET(contextPath+"/cronjob/:id", controllers.CronjobShow)
	r.PUT(contextPath+"/cronjob/:id", controllers.CronjobUpdate)
	r.DELETE(contextPath+"/cronjob/:id", controllers.CronjobDelete)

	r.GET(contextPath+"/schedule", controllers.CronjobNextOccurIndex)
	r.GET(contextPath+"/schedule/:id", controllers.CronjobNextOccurShow)
	r.DELETE(contextPath+"/schedule/:id", controllers.CronjobNextOccurDelete)

	r.Run()
}
