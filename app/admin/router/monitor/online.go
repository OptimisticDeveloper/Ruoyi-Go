package monitor

import (
	"ruoyi-go/app/admin/api/monitor"
	"ruoyi-go/utils"

	"github.com/gin-gonic/gin"
)

func InitOnLine(e *gin.Engine) {
	// 在线用户相关
	v := e.Group("monitor/online")
	{
		auth := v.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.GET("list", monitor.ListOnLine)
			auth.DELETE("/:tokenId", monitor.DetectOnLine)
		}
	}
}
