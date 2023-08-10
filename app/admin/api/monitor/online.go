package monitor

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ruoyi-go/app/admin/model/monitor"
	"ruoyi-go/utils"
	"ruoyi-go/utils/R"
)

func ListOnLine(context *gin.Context) {

	context.JSON(http.StatusOK, gin.H{
		"msg":  "操作成功",
		"code": http.StatusOK,
		"data": "",
	})
}

func DetectOnLine(context *gin.Context) {
	var tokenId = context.Param("tokenId")
	var result = DelectOnLines(utils.Split(tokenId))
	context.JSON(http.StatusOK, result)
}

func DelectOnLines(operId []int) R.Result {
	if err := utils.MysqlDb.Model(&monitor.SysOperLog{}).
		Delete("oper_id in ?", operId).Error; err != nil {
		return R.ReturnFailMsg(err.Error())
	}
	return R.ReturnSuccess("操作成功")
}
