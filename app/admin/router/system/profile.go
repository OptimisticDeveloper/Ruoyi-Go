package system

import (
	"ruoyi-go/app/admin/api/system"
	"ruoyi-go/utils"

	"github.com/gin-gonic/gin"
)

func InitProfile(e *gin.Engine) {
	v := e.Group("system/user")
	{
		auth := v.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.PUT("profile/updatePwd", system.UpdatePwdHandler)
			auth.GET("profile", system.ProfileHandler)
			auth.PUT("profile", system.PostProfileHandler)
			auth.POST("profile/avatar", system.AvatarHandler)
		}
	}
}
