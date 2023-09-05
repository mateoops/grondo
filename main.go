package main

import (
	"fmt"
	"grondo/controllers"
	sql "grondo/db"
	"grondo/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	utils.LoadYamlConfig()
	sql.InitDB()
}

func main() {
	fmt.Println("Config:")
	fmt.Println(utils.AppConfig)
	contextPath := utils.AppConfig.Server.ContextPath
	apiVersion := "v1"
	apiLocation := fmt.Sprintf("%s/%s", contextPath, apiVersion)

	r := gin.Default()

	r.POST(apiLocation+"/cronjob", controllers.CronjobCreate)
	r.GET(apiLocation+"/cronjob", controllers.CronjobIndex)
	r.GET(apiLocation+"/cronjob/:id", controllers.CronjobShow)
	r.PUT(apiLocation+"/cronjob/:id", controllers.CronjobUpdate)
	r.DELETE(apiLocation+"/cronjob/:id", controllers.CronjobDelete)

	r.Run()
}
