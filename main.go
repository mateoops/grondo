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

	r := gin.Default()

	r.POST("/api/v1/cronjob", controllers.CronjobCreate)
	r.GET("/api/v1/cronjob", controllers.CronjobIndex)
	r.GET("/api/v1/cronjob/:id", controllers.CronjobShow)
	r.PUT("/api/v1/cronjob/:id", controllers.CronjobUpdate)
	r.DELETE("/api/v1/cronjob/:id", controllers.CronjobDelete)

	r.Run()
}
