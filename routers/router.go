package routers

import (
	"fmt"
	"grondo/controllers"
	"grondo/utils/config"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	contextPath := fmt.Sprintf("%s/%s", config.AppConfig.Server.ContextPath, config.AppConfig.Server.ApiVersion)

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
	r.GET(contextPath+"/queue", controllers.CronjobQueueIndex)
	r.GET(contextPath+"/queue/:id", controllers.CronjobQueueShow)

	// JobExecLog routes
	r.GET(contextPath+"/log", controllers.JobExecLogIndex)
	r.GET(contextPath+"/log/:id", controllers.JobExecLogShow)

	// Machine routes
	r.POST(contextPath+"/machine", controllers.MachineCreate)
	r.GET(contextPath+"/machine", controllers.MachineIndex)
	r.GET(contextPath+"/machine/:id", controllers.MachineShow)
	r.PUT(contextPath+"/machine/:id", controllers.MachineUpdate)
	r.DELETE(contextPath+"/machine/:id", controllers.MachineDelete)

	// UserWithPassword routes
	r.POST(contextPath+"/user/password", controllers.UserPassCreate)
	r.GET(contextPath+"/user/password", controllers.UserPassIndex)
	r.GET(contextPath+"/user/password/:id", controllers.UserPassShow)
	r.PUT(contextPath+"/user/password/:id", controllers.UserPassUpdate)
	r.DELETE(contextPath+"/user/password/:id", controllers.UserPassDelete)

	// UserWithSSHKey routes
	r.POST(contextPath+"/user/sshkey", controllers.UserSSHCreate)
	r.GET(contextPath+"/user/sshkey", controllers.UserSSHIndex)
	r.GET(contextPath+"/user/sshkey/:id", controllers.UserSSHShow)
	r.PUT(contextPath+"/user/sshkey/:id", controllers.UserSSHUpdate)
	r.DELETE(contextPath+"/user/sshkey/:id", controllers.UserSSHDelete)

	return r
}
