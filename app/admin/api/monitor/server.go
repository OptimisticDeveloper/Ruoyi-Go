package monitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ruoyi-go/app/admin/model/monitor"
)

func ServerData(context *gin.Context) {
	var operLog = monitor.SysOperLog{
		Title:        "操作日志",
		BusinessType: monitor.BusinessTypeAdd,
		Method:       "GET",
		OperatorType: monitor.OperatorTypeAdmin,
		OperParam:    "",
		JsonResult:   "",
		Status:       "0",
		ErrorMsg:     "",
	}
	operLog.OperationLogAdd(context)
	var server = monitor.GetServerInfo(context)
	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": server,
	})
}
