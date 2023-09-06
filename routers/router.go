package routers

import (
	"fmt"
	"grondo/controllers"
	"grondo/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	contextPath := fmt.Sprintf("%s/%s", utils.AppConfig.Server.ContextPath, utils.AppConfig.Server.ApiVersion)

	r := gin.Default()

	// Cronjob routes
	r.POST(contextPath+"/cronjob", controllers.CronjobCreate)
	r.GET(contextPath+"/cronjob", controllers.CronjobIndex)
	r.GET(contextPath+"/cronjob/:id", controllers.CronjobShow)
	r.PUT(contextPath+"/cronjob/:id", controllers.CronjobUpdate)
	r.DELETE(contextPath+"/cronjob/:id", controllers.CronjobDelete)

	// Schedule routes
	r.GET(contextPath+"/schedule", controllers.ScheduleIndex)
	r.GET(contextPath+"/schedule/:id", controllers.ScheduleShow)

	// Queue routes
	r.GET(contextPath+"/queue", controllers.CronjobQueue)
	r.GET(contextPath+"/queue/:id", controllers.CronjobQueueShow)
	return r
}
