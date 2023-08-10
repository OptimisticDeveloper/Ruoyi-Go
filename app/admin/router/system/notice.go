package system

import (
	"ruoyi-go/app/admin/api/system"
	"ruoyi-go/utils"

	"github.com/gin-gonic/gin"
)

func InitNotice(e *gin.Engine) {
	// 消息相关
	v := e.Group("system/notice")
	{
		auth := v.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.GET("/list", system.ListNotice)
			auth.GET("/:noticeId", system.GetNotice)
			auth.POST("", system.SaveNotice)
			auth.PUT("", system.UploadNotice)
			auth.DELETE("/:noticeIds", system.DeleteNotice)
		}
	}
}
