package html

import (
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"ruoyi-go/app/html/web"
	"time"
)

func Routers(e *gin.Engine) {
	memoryStore := persist.NewMemoryStore(1 * time.Minute)
	handlerFunc := cache.CacheByRequestURI(memoryStore, 2*time.Second)
	e.GET("/", web.IndexHandler)
	e.GET("/admin", web.IndexAdminHandler)
	e.GET("/old", web.IndexOldHandler)
	e.GET("/protocol.html", handlerFunc, web.ProtocolHandler)
}
