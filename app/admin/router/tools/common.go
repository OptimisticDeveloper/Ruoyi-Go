package tools

import (
	"ruoyi-go/app/admin/api/tools"
	"ruoyi-go/utils"

	"github.com/gin-gonic/gin"
)

func InitCommon(e *gin.Engine) {
	// 公共相关
	v := e.Group("common")
	{
		auth := v.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.GET("/download", tools.GetDownload)
			auth.POST("/upload", tools.UploadCommon)
			auth.POST("/uploads", tools.UploadCommons)
			auth.GET("/download/resource", tools.UploadRsource)
		}
	}
}
