package monitor

import (
	"ruoyi-go/app/admin/api/monitor"
	"ruoyi-go/utils"

	"github.com/gin-gonic/gin"
)

func InitJob(e *gin.Engine) {
	// 定时任务相关
	v1 := e.Group("monitor/job")
	{
		auth := v1.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.GET("list", monitor.ListJob)
			auth.POST("export", monitor.ExportJob)
			auth.GET(":jobId", monitor.GetJobById)
			auth.POST("", monitor.SaveJob)
			auth.PUT("", monitor.UploadJob)
			auth.PUT("changeStatus", monitor.ChangeStatus)
			auth.PUT("run", monitor.RunJob)
			auth.DELETE(":jobIds", monitor.DelectJob)
		}
	}
}

func InitJobLog(e *gin.Engine) {
	// 路由权限相关
	v2 := e.Group("monitor/jobLog")
	{
		auth := v2.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			// 需要权限
			auth.GET("list", monitor.ListJobLog)
			auth.POST("export", monitor.ExportJobLog)
			auth.GET(":configId", monitor.GetJobLog)
			auth.DELETE(":jobLogIds", monitor.DetectJobLog)
			auth.DELETE("clean", monitor.ClearJobLog)
		}
	}
}
