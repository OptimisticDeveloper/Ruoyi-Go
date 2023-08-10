package monitor

import (
	"ruoyi-go/app/admin/api/monitor"
	"ruoyi-go/utils"

	"github.com/gin-gonic/gin"
)

// InitLogininfor 登录日志
func InitLogininfor(e *gin.Engine) {
	// 登录日志相关
	v := e.Group("monitor/logininfor")
	{
		auth := v.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.GET("/list", monitor.LoginInformListHandler)
			auth.POST("/export", monitor.ExportHandler)
			auth.DELETE("/:infoIds", monitor.DeleteByIdHandler)
			auth.DELETE("/clean", monitor.CleanHandler)
			auth.GET("/unlock/:userName", monitor.UnlockHandler)
		}
	}
}
