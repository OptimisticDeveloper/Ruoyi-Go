package system

import (
	"ruoyi-go/app/admin/api/system"
	"ruoyi-go/utils"

	"github.com/gin-gonic/gin"
)

func InitDept(e *gin.Engine) {
	// 部门相关
	v := e.Group("system/dept")
	{
		auth := v.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			// 部门数据
			auth.GET("/list", system.ListDept)
			auth.GET("/list/exclude/:deptId", system.ExcludeDept)
			auth.GET("/:deptId", system.GetDept)
			auth.POST("", system.SaveDept)
			auth.PUT("", system.UpDataDept)
			auth.DELETE("/:deptId", system.DeleteDept)
		}
	}
}
