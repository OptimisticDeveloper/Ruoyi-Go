package admin

import (
	api "ruoyi-go/app/admin/api/system"
	"ruoyi-go/app/admin/router/monitor"
	"ruoyi-go/app/admin/router/system"
	"ruoyi-go/app/admin/router/tools"
	"ruoyi-go/utils"
	"time"

	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {

	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	handlerFunc := cache.CacheByRequestURI(memoryStore, 2*time.Second)

	e.GET("/index", api.IndexHandler)
	// 登录
	e.GET("/captchaImage", api.CaptchaImageHandler)
	e.POST("/login", api.LoginHandler)

	v1 := e.Group("/")
	{
		auth := v1.Group("")
		auth.Use(utils.JWTAuthMiddleware())
		{
			auth.GET("getInfo", handlerFunc, api.GetInfoHandler)
			auth.POST("logout", api.LogoutHandler)
			/*获取用户授权菜单*/
			auth.GET("getRouters", handlerFunc, api.GetRoutersHandler)
		}
	}
	/*system*/
	system.InitProfile(e)
	system.InitDict(e)
	system.InitUser(e)
	system.InitMenu(e)
	system.InitPost(e)
	system.InitNotice(e)
	system.InitRole(e)
	system.InitConfig(e)
	system.InitDept(e)

	/*monitor*/
	monitor.InitCache(e)
	monitor.InitLogininfor(e)
	monitor.InitJob(e)
	monitor.InitJobLog(e)
	monitor.InitOnLine(e)
	monitor.InitOperlog(e)
	monitor.InitServer(e)

	/*tools*/
	tools.InitCommon(e)

}
