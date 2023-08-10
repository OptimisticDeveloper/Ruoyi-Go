package monitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ruoyi-go/app/admin/model/monitor"
)

func CacheHandler(context *gin.Context) {
	var list []monitor.SysCache
	list = append(list, monitor.SysCache{
		CacheName: "login_tokens:",
		Remark:    "用户信息",
	})
	list = append(list, monitor.SysCache{
		CacheName: "sys_config:",
		Remark:    "配置信息",
	})
	list = append(list, monitor.SysCache{
		CacheName: "sys_dict:",
		Remark:    "数据字典",
	})
	list = append(list, monitor.SysCache{
		CacheName: "captcha_codes:",
		Remark:    "验证码",
	})
	list = append(list, monitor.SysCache{
		CacheName: "repeat_submit:",
		Remark:    "防重提交",
	})
	list = append(list, monitor.SysCache{
		CacheName: "rate_limit:",
		Remark:    "限流处理",
	})
	list = append(list, monitor.SysCache{
		CacheName: "pwd_err_cnt:",
		Remark:    "密码错误次数",
	})
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": list,
	})
}

func GetCacheKeysHandler(context *gin.Context) {
	cacheName := context.Param("cacheName")
	var list = []string{
		"" + cacheName + ":123456",
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": list,
	})
}

func GetCacheValueHandler(context *gin.Context) {
	cacheName := context.Param("cacheName")
	cacheKey := context.Param("cacheKey")
	var cache = monitor.SysCache{
		CacheName:  cacheName,
		CacheKey:   cacheKey,
		CacheValue: "123456",
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": cache,
	})
}

func ClearCacheNameHandler(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
	})
}

func ClearCacheKeyHandler(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
	})
}

func ClearCacheAllHandler(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
	})
}
