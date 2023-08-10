package routers

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/admin"
	"ruoyi-go/app/html"
	"ruoyi-go/config"
	"ruoyi-go/utils"
)

// Init 初始化
func Init() *gin.Engine {
	gin.SetMode(config.Server.RunMode)
	r := gin.New()

	r.Use(utils.Logger())
	r.Use(gin.Recovery())
	/*自定义错误*/
	r.Use(utils.Recover)

	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.Use(utils.Core())
	r.HTMLRender = createRender()

	r.Static("/profile", "./static/images")
	r.Static("/admin", "./view/admin")
	r.Static("/static", "./view/mobile/static")
	r.Static("/favicon.ico", "./view/admin/static/favicon.ico")

	// 加载多个APP的路由配置
	html.Routers(r)
	admin.Routers(r)

	return r
}

// 不同模板设置
func createRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "view/admin/index.html")
	p.AddFromFiles("mobile", "view/mobile/index.html")
	p.AddFromFiles("mobile_old", "view/mobile_old/index.html")
	p.AddFromFiles("protocol", "view/template/protocol.tmpl")
	return p
}
