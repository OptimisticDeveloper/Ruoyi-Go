package monitor

import (
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/admin/api/monitor"
	"ruoyi-go/utils"
)

func InitServer(e *gin.Engine) {
	// 服务监控
	v := e.Group("monitor")
	{
		auth := v.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.GET("/server", monitor.ServerData)
		}
	}
}
