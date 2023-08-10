package system

import (
	"ruoyi-go/app/admin/api/system"
	"ruoyi-go/utils"

	"github.com/gin-gonic/gin"
)

func InitMenu(e *gin.Engine) {
	// 菜单相关
	v := e.Group("system/menu")
	{
		auth := v.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.GET("/list", system.ListMenu)
			auth.GET("/:menuId", system.GetMenuInfo)
			auth.GET("/treeselect", system.GetTreeSelect)
			auth.GET("/roleMenuTreeselect/:roleId", system.TreeSelectByRole)
			auth.POST("", system.SaveMenu)
			auth.PUT("", system.UploadMenu)
			auth.DELETE("/:menuId", system.DeleteMenu)
		}
	}
}
