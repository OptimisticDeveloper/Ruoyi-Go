package system

import (
	"ruoyi-go/app/admin/api/system"
	"ruoyi-go/utils"

	"github.com/gin-gonic/gin"
)

func InitPost(e *gin.Engine) {
	v := e.Group("system/post")
	{
		auth := v.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.GET("/list", system.ListPost)
			auth.POST("/export", system.ExportPost)
			auth.GET("/:postId", system.GetPostInfo)
			auth.POST("", system.SavePost)
			auth.PUT("", system.UploadPost)
			auth.DELETE("/:postIds", system.DeletePost)
			auth.GET("/optionselect", system.GetPostOptionSelect)
		}
	}
}
