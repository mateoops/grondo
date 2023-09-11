package routers

import (
	"fmt"
	"grondo/controllers"
	_ "grondo/docs"
	"grondo/utils/config"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	v1 := r.Group(fmt.Sprintf("%s/%s", config.AppConfig.Server.ContextPath, config.AppConfig.Server.ApiVersion))
	{
		cronjob := v1.Group("/cronjob")
		{
			cronjob.POST("", controllers.CronjobCreate)
			cronjob.GET("", controllers.CronjobIndex)
			cronjob.GET(":id", controllers.CronjobShow)
			cronjob.PUT(":id", controllers.CronjobUpdate)
			cronjob.DELETE(":id", controllers.CronjobDelete)
		}
		schedule := v1.Group("/schedule")
		{
			schedule.GET("", controllers.ScheduleIndex)
			schedule.GET(":id", controllers.ScheduleShow)
		}
		queue := v1.Group("/queue")
		{
			queue.GET("", controllers.CronjobQueueIndex)
			queue.GET(":id", controllers.CronjobQueueShow)
		}
		log := v1.Group("/log")
		{
			log.GET("", controllers.JobExecLogIndex)
			log.GET(":id", controllers.JobExecLogShow)
		}
		machine := v1.Group("/machine")
		{
			machine.POST("", controllers.MachineCreate)
			machine.GET("", controllers.MachineIndex)
			machine.GET(":id", controllers.MachineShow)
			machine.PUT(":id", controllers.MachineUpdate)
			machine.DELETE(":id", controllers.MachineDelete)
		}
		user := v1.Group("/user")
		{
			userWithPassword := user.Group("/password")
			{
				userWithPassword.POST("", controllers.UserPassCreate)
				userWithPassword.GET("", controllers.UserPassIndex)
				userWithPassword.GET(":id", controllers.UserPassShow)
				userWithPassword.PUT(":id", controllers.UserPassUpdate)
				userWithPassword.DELETE(":id", controllers.UserPassDelete)
			}
			userWithSSHKey := user.Group("/sshkey")
			{
				userWithSSHKey.POST("", controllers.UserSSHCreate)
				userWithSSHKey.GET("", controllers.UserSSHIndex)
				userWithSSHKey.GET(":id", controllers.UserSSHShow)
				userWithSSHKey.PUT(":id", controllers.UserSSHUpdate)
				userWithSSHKey.DELETE(":id", controllers.UserSSHDelete)
			}
		}
		v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	return r
}
